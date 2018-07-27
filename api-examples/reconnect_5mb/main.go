package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_url]
	// Set reconnect buffer size in bytes (5 MB)
	nc, err := nats.Connect("nats://demo.nats.io:4222", nats.ReconnectBufSize(5*1024*1024))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_url]
}
