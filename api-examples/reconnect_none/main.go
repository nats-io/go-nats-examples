package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin reconnect_none]
	// Disable reconnect attempts
	nc, err := nats.Connect("demo.nats.io", nats.NoReconnect())
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end reconnect_none]
}
