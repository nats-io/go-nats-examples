package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_multiple]
	nc, err := nats.Connect("nats://localhost:1222, nats://localhost:1223, nats://localhost:1224")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection
	// [end connect_multiple]
}
