package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin connect_nkey]
	opt, err := nats.NkeyOptionFromSeed("seed.txt")
	if err != nil {
		log.Fatal(err)
	}
	nc, err := nats.Connect("127.0.0.1", opt)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_nkey]
}
