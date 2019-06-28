package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin slow_listener]
	// Set the callback that will be invoked when an asynchronous error occurs.
	nc, err := nats.Connect("demo.nats.io", nats.ErrorHandler(logSlowConsumer))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end slow_listener]
}

func logSlowConsumer(nc *nats.Conn, sub *nats.Subscription, err error) {
	if err == nats.ErrSlowConsumer {
		dropped, _ := sub.Dropped()
		log.Printf("Slow consumer on subject %q: dropped %d messages\n",
			sub.Subject, dropped)
	}
}
