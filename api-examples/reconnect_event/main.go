package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin reconnect_event]
	// Connection event handlers are invoked asynchronously
	// and the state of the connection may have changed when
	// the callback is invoked.
	nc, err := nats.Connect("nats://demo.nats.io:4222",
		nats.DisconnectHandler(func(nc *nats.Conn) {
			// handle disconnect event
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			// handle reconnect event
		}))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end reconnect_event]
}
