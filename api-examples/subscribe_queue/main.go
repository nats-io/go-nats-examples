package main

import (
	"log"
	"sync"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin subscribe_queue]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Use a WaitGroup to wait for 10 messages to arrive
	wg := sync.WaitGroup{}
	wg.Add(10)

	// Create a queue subscription on "updates" with queue name "workers"
	if _, err := nc.QueueSubscribe("updates", "worker", func(m *nats.Msg) {
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for messages to come in
	wg.Wait()

	// Close the connection
	nc.Close()
	// [end subscribe_queue]
}
