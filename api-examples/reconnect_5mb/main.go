package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin reconnect_5mb]
	// Set reconnect buffer size in bytes (5 MB)
	nc, err := nats.Connect("demo.nats.io", nats.ReconnectBufSize(5*1024*1024))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end reconnect_5mb]
}
