package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin flush]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	if err := nc.Publish("updates", []byte("All is Well")); err != nil {
		log.Fatal(err)
	}
	// Sends a PING and wait for a PONG from the server, up to the given timeout.
	// This gives guarantee that the server has processed above message.
	if err := nc.FlushTimeout(time.Second); err != nil {
		log.Fatal(err)
	}
	// [end flush]
}
