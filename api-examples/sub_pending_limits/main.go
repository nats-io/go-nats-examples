package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin slow_pending_limits]
	nc, err := nats.Connect("nats://demo.nats.io:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe
	sub1, err := nc.Subscribe("updates", func(m *nats.Msg) {})
	if err != nil {
		log.Fatal(err)
	}

	// Set limits of 1000 messages or 5MB, whichever comes first
	sub1.SetPendingLimits(1000, 5*1024*1024)

	// Subscribe
	sub2, err := nc.Subscribe("updates", func(m *nats.Msg) {})
	if err != nil {
		log.Fatal(err)
	}

	// Set no limits for this subscription
	sub2.SetPendingLimits(-1, -1)

	// Close the connection
	nc.Close()
	// [end slow_pending_limits]
}
