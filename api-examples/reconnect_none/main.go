package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_url]
	// Disable reconnect attempts
	nc, err := nats.Connect("nats://demo.nats.io:4222", nats.NoReconnect())
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_url]
}
