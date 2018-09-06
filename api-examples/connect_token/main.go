package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_token]
	// Set a token
	nc, err := nats.Connect("localhost", nats.Token("mytoken"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_token]
}
