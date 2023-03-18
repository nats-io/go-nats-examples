// Copyright 2012-2019 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/tls"
	"flag"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

// NOTE: Can test with demo servers.
// nats-req -s demo.nats.io <subject> <msg>
// nats-req -s demo.nats.io:4443 <subject> <msg> (TLS version)

func usage() {
	log.Printf("Usage: nats-req [-s server] [-i] [-creds file] <subject> <msg>\n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var userCreds = flag.String("creds", "", "User Credentials File")
	var showHelp = flag.Bool("h", false, "Show help message")
	var insecure = flag.Bool("i", false, "If TLS set, don't verify certs")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	args := flag.Args()
	if len(args) < 2 {
		showUsageAndExit(1)
	}

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Sample Requestor")}

	if *insecure {
		opts = append(opts, nats.Secure(&tls.Config{
			InsecureSkipVerify: true,
		}))
	}

	// Use UserCredentials
	if *userCreds != "" {
		opts = append(opts, nats.UserCredentials(*userCreds))
	}

	if *insecure {
		opts = append(opts, nats.Secure(&tls.Config{
			InsecureSkipVerify: true,
		}))
	}

	// Connect to NATS
	nc, err := nats.Connect(*urls, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	subj, payload := args[0], []byte(args[1])

	msg, err := nc.Request(subj, payload, 2*time.Second)
	if err != nil {
		if nc.LastError() != nil {
			log.Fatalf("%v for request", nc.LastError())
		}
		log.Fatalf("%v for request", err)
	}

	log.Printf("Published [%s] : '%s'", subj, payload)
	log.Printf("Received  [%v] : '%s'", msg.Subject, string(msg.Data))
}
