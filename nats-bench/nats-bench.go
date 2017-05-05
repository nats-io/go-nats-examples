// Copyright 2015 Apcera Inc. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats/bench"
	"math/rand"
)

// Some sane defaults
const (
	DefaultNumMsgs     = 100000
	DefaultNumPubs     = 1
	DefaultNumSubs     = 0
	DefaultMessageSize = 128
)

func usage() {
	log.Fatalf("Usage: nats-bench [-s server (%s)] [--tls] [-np NUM_PUBLISHERS] [-ns NUM_SUBSCRIBERS] [-n NUM_MSGS] [-ms MESSAGE_SIZE] [-csv csvfile] <subject>\n", nats.DefaultURL)
}

var benchmark *bench.Benchmark

func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var tls = flag.Bool("tls", false, "Use TLS Secure Connection")
	var numPubs = flag.Int("np", DefaultNumPubs, "Number of Concurrent Publishers")
	var numSubs = flag.Int("ns", DefaultNumSubs, "Number of Concurrent Subscribers")
	var numMsgs = flag.Int("n", DefaultNumMsgs, "Number of Messages to Publish")
	var msgSize = flag.Int("ms", DefaultMessageSize, "Size of the message.")
	var csvFile = flag.String("csv", "", "Save bench data to csv file")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		usage()
	}

	if *numMsgs <= 0 {
		log.Fatal("Number of messages should be greater than zero.")
	}

	// Setup the option block
	opts := nats.DefaultOptions
	opts.Servers = strings.Split(*urls, ",")
	for i, s := range opts.Servers {
		opts.Servers[i] = strings.Trim(s, " ")
	}
	opts.Secure = *tls

	benchmark = bench.NewBenchmark("NATS", *numSubs, *numPubs)

	var startwg sync.WaitGroup
	var donewg sync.WaitGroup

	donewg.Add(*numPubs + *numSubs)

	//opts.Timeout = 30 * time.Second
	//opts.AllowReconnect = true
	//opts.MaxReconnect = 5

	rand.Seed(time.Now().Unix())

	// Run Subscribers first
	startwg.Add(*numSubs)
	for i := 0; i < *numSubs; i++ {
		// add a random wait time
		time.Sleep(time.Duration(rand.Int31n(10)) * time.Millisecond)
		go runSubscriber(&startwg, &donewg, opts, *numMsgs, *msgSize)
	}
	startwg.Wait()

	// Now Publishers
	startwg.Add(*numPubs)
	pubCounts := bench.MsgsPerClient(*numMsgs, *numPubs)
	for i := 0; i < *numPubs; i++ {
		go runPublisher(&startwg, &donewg, opts, pubCounts[i], *msgSize)
	}

	log.Printf("Starting benchmark [msgs=%d, msgsize=%d, pubs=%d, subs=%d]\n", *numMsgs, *msgSize, *numPubs, *numSubs)

	startwg.Wait()
	donewg.Wait()

	benchmark.Close()

	fmt.Print(benchmark.Report())

	if len(*csvFile) > 0 {
		csv := benchmark.CSV()
		if err := ioutil.WriteFile(*csvFile, []byte(csv), 0644); err != nil {
			log.Printf("Unable to save metric data to file %s: %v", *csvFile, err)
		}
		fmt.Printf("Saved metric data in csv file %s\n", *csvFile)
	}
}

func runPublisher(startwg, donewg *sync.WaitGroup, opts nats.Options, numMsgs int, msgSize int) {
	nc, err := opts.Connect()
	if err != nil {
		log.Printf("Publisher can't connect: %v\n", err)
		donewg.Done()
		return
	}
	defer nc.Close()
	startwg.Done()

	args := flag.Args()
	subj := args[0]
	var msg []byte
	if msgSize > 0 {
		msg = make([]byte, msgSize)
	}

	start := time.Now()

	rand.Seed(start.Unix())

	for i := 0; i < numMsgs; i++ {

		// add a random wait time
		time.Sleep(time.Duration(rand.Int31n(10)) * time.Millisecond)
		nc.Publish(subj, msg)
	}

	benchmark.AddPubSample(bench.NewSample(numMsgs, msgSize, start, time.Now(), nc))

	if err = nc.Flush(); err != nil {
		log.Printf("Unable to flush: %v", err)
	}

	donewg.Done()
}

func runSubscriber(startwg, donewg *sync.WaitGroup, opts nats.Options, numMsgs int, msgSize int) {
	received := 0
	start := time.Now()

	opts.ClosedCB = func(nc *nats.Conn) {
		benchmark.AddSubSample(bench.NewSample(numMsgs, msgSize, start, time.Now(), nc))
		donewg.Done()
	}

	opts.DisconnectedCB = func(nc *nats.Conn) {
		if !nc.IsClosed() {
			nc.Close()
		}
	}

	nc, err := opts.Connect()

	if err != nil {
		log.Printf("Can't connect: %v\n", err)
		donewg.Done()
		startwg.Done()
		return
	}

	args := flag.Args()
	subj := args[0]

	_, err = nc.Subscribe(subj, func(msg *nats.Msg) {
		if received == 0 {
			start = time.Now()
		}
		received++
		if received >= numMsgs {
			nc.Close()
		}
	})

	if err != nil {
		donewg.Done()
		log.Printf("Can't subscribe to %q: %v", subj, err)

	}
	if err = nc.Flush(); err != nil {
		donewg.Done()
		log.Printf("Error flushing: %v", err)
	}
	startwg.Done()
}
