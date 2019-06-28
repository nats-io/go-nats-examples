package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin connect_pedantic]
	opts := nats.GetDefaultOptions()
	opts.Url = "demo.nats.io"
	// Turn on Pedantic
	opts.Pedantic = true
	nc, err := opts.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_pedantic]
}
