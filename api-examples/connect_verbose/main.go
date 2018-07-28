package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_verbose]
	opts := nats.GetDefaultOptions()
	opts.Url = "nats://demo.nats.io:4222"
	// Turn on Verbose
	opts.Verbose = true
	nc, err := opts.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_verbose]
}
