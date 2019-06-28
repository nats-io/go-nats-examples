package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin drain_sub]

	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	done := sync.WaitGroup{}
	done.Add(1)

	count := 0
	errCh := make(chan error, 1)

	msgAfterDrain := "not this one"

	// Just to not collide using the demo server with other users.
	subject := nats.NewInbox()

	// This callback will process each message slowly
	sub, err := nc.Subscribe(subject, func(m *nats.Msg) {
		if string(m.Data) == msgAfterDrain {
			errCh <- fmt.Errorf("Should not have received this message")
			return
		}
		time.Sleep(100 * time.Millisecond)
		count++
		if count == 2 {
			done.Done()
		}
	})

	// Send 2 messages
	for i := 0; i < 2; i++ {
		nc.Publish(subject, []byte("hello"))
	}

	// Call Drain on the subscription. It unsubscribes but
	// wait for all pending messages to be processed.
	if err := sub.Drain(); err != nil {
		log.Fatal(err)
	}

	// Send one more message, this message should not be received
	nc.Publish(subject, []byte(msgAfterDrain))

	// Wait for the subscription to have processed the 2 messages.
	done.Wait()

	// Now check that the 3rd message was not received
	select {
	case e := <-errCh:
		log.Fatal(e)
	case <-time.After(200 * time.Millisecond):
		// OK!
	}

	// [end drain_sub]
}
