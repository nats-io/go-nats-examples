package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_url]
	// If connecting to the default port, the URL can be simplified
	// to just the hostname/IP.
	// That is, the connect below is equivalent to:
	// nats.Connect("nats://demo.nats.io:4222")
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_url]
}
