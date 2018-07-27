package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_url]
	// Set max reconnects attempts
	nc, err := nats.Connect("nats://demo.nats.io:4222", nats.MaxReconnects(10))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_url]
}
