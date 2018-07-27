package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_url]
	// Set reconnect interval to 10 seconds
	nc, err := nats.Connect("nats://demo.nats.io:4222", nats.ReconnectWait(10*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_url]
}
