package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin no_echo]
	// Turn off echo
	nc, err := nats.Connect("nats://demo.nats.io:4222", nats.NoEcho())
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end no_echo]
}
