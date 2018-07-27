package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_options]
	nc, err := nats.Connect(nats.DefaultURL, nats.Timeout(10*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_options]
}
