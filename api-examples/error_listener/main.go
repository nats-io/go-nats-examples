package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin error_listener]
	// Set the callback that will be invoked when an asynchronous error occurs.
	nc, err := nats.Connect("demo.nats.io",
		nats.ErrorHandler(func(_ *nats.Conn, _ *nats.Subscription, err error) {
			log.Printf("Error: %v", err)
		}))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end error_listener]
}
