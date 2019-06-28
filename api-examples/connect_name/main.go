package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin connect_name]
	// Set a connection name
	nc, err := nats.Connect("demo.nats.io", nats.Name("API Name Example"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_name]
}
