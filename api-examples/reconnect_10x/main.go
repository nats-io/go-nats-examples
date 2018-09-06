package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin reconnect_10x]
	// Set max reconnects attempts
	nc, err := nats.Connect("demo.nats.io", nats.MaxReconnects(10))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end reconnect_10x]
}
