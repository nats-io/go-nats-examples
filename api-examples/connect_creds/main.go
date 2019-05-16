package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_creds]
	opt, err := nats.NkeyOptionFromSeed("seed.txt")
	if err != nil {
		log.Fatal(err)
	}
	nc, err := nats.Connect("localhost", opt)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_creds]
}
