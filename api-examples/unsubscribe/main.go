package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin unsubscribe]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Sync Subscription
	sub, err := nc.SubscribeSync("updates")
	if err != nil {
		log.Fatal(err)
	}
	if err := sub.Unsubscribe(); err != nil {
		log.Fatal(err)
	}

	// Async Subscription
	sub, err = nc.Subscribe("updates", func(_ *nats.Msg) {})
	if err != nil {
		log.Fatal(err)
	}
	if err := sub.Unsubscribe(); err != nil {
		log.Fatal(err)
	}

	// [end unsubscribe]
}
