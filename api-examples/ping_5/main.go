package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin ping_5]

	// Set maximum number of PINGs out without getting a PONG back
	// before the connection will be disconnected as a stale connection.
	nc, err := nats.Connect("demo.nats.io", nats.Name("API MaxPing Example"), nats.MaxPingsOutstanding(5))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end ping_5]
}
