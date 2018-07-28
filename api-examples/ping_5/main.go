package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin ping_5]
	opts := nats.GetDefaultOptions()
	opts.Url = "nats://demo.nats.io:4222"
	// Set maximum number of PINGs out without getting a PONG back
	// before the connection will be disconnected as a stale connection.
	opts.MaxPingsOut = 5

	nc, err := opts.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end ping_5]
}
