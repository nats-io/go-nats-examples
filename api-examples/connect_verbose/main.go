package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin connect_verbose]
	opts := nats.GetDefaultOptions()
	opts.Url = "demo.nats.io"
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
