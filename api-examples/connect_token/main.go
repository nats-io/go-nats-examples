package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin connect_token]
	// Set a token
	nc, err := nats.Connect("127.0.0.1", nats.Name("API Token Example"), nats.Token("mytoken"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_token]
}
