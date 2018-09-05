package main

import (
	"log"
	"sync"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin drain_conn]
	wg := sync.WaitGroup{}
	wg.Add(1)

	errCh := make(chan error, 1)

	nc, err := nats.Connect("nats://demo.nats.io:4222",
		nats.DrainTimeout(10*time.Second),
		nats.ErrorHandler(func(_ *nats.Conn, _ *nats.Subscription, err error) {
			errCh <- err
		}),
		nats.ClosedHandler(func(_ *nats.Conn) {
			wg.Done()
		}))
	if err != nil {
		log.Fatal(err)
	}
	// Do something with the connection

	// Drain the connection, which will close it when done.
	if err := nc.Drain(); err != nil {
		log.Fatal(err)
	}

	// Wait for the connection to be closed.
	wg.Wait()

	// Check if there was an error
	select {
	case e := <-errCh:
		log.Fatal(e)
	default:
	}

	// [end drain_conn]
}
