package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin unsubscribe_auto]
	nc, err := nats.Connect("nats://demo.nats.io:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Sync Subscription
	sub, err := nc.SubscribeSync("updates")
	if err != nil {
		log.Fatal(err)
	}
	if err := sub.AutoUnsubscribe(1); err != nil {
		log.Fatal(err)
	}

	// Async Subscription
	sub, err = nc.Subscribe("updates", func(_ *nats.Msg) {})
	if err != nil {
		log.Fatal(err)
	}
	if err := sub.AutoUnsubscribe(1); err != nil {
		log.Fatal(err)
	}

	// Close the connection
	nc.Close()
	// [end unsubscribe_auto]
}
