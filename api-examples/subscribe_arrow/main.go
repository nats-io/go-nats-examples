package main

import (
	"log"
	"sync"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin subscribe_arrow]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Use a WaitGroup to wait for 4 messages to arrive
	wg := sync.WaitGroup{}
	wg.Add(4)

	// Subscribe
	if _, err := nc.Subscribe("time.>", func(m *nats.Msg) {
		log.Printf("%s: %s", m.Subject, m.Data)
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for the 4 messages to come in
	wg.Wait()

	// Close the connection
	nc.Close()
	// [end subscribe_arrow]
}
