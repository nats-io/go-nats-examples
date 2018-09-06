package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_name]
	// Set a connection name
	nc, err := nats.Connect("demo.nats.io", nats.Name("my-connection"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_name]
}
