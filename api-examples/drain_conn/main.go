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

	// To simulate a timeout, you would set the DrainTimeout()
	// to a value less than the time spent in the message callback,
	// so say: nats.DrainTimeout(10*time.Millisecond).

	nc, err := nats.Connect("demo.nats.io",
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

	// Subscribe, but add some delay while processing.
	if _, err := nc.Subscribe("foo", func(_ *nats.Msg) {
		time.Sleep(200 * time.Millisecond)
	}); err != nil {
		log.Fatal(err)
	}

	// Publish a message
	if err := nc.Publish("foo", []byte("hello")); err != nil {
		log.Fatal(err)
	}

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
