package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin error_listener]
	// Set the callback that will be invoked when an asynchronous error occurs.
	nc, err := nats.Connect("demo.nats.io",
		nats.ErrorHandler(func(nc *nats.Conn, sub *nats.Subscription, err error) {
			log.Printf("Error: %v", err)
		}))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end error_listener]
}
