package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin reconnect_10s]
	// Set reconnect interval to 10 seconds
	nc, err := nats.Connect("demo.nats.io", nats.ReconnectWait(10*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end reconnect_10s]
}
