package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin flush]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Just to not collide using the demo server with other users.
	subject := nats.NewInbox()

	if err := nc.Publish(subject, []byte("All is Well")); err != nil {
		log.Fatal(err)
	}
	// Sends a PING and wait for a PONG from the server, up to the given timeout.
	// This gives guarantee that the server has processed the above message.
	if err := nc.FlushTimeout(time.Second); err != nil {
		log.Fatal(err)
	}
	// [end flush]
}
