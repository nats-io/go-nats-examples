package main

import (
	"log"
	"sync"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin subscribe_async]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	if _, err := nc.Subscribe("updates", func(m *nats.Msg) {
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()

	// Close the connection
	nc.Close()
	// [end subscribe_async]
}
