package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin wildcard_tester]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	zoneID, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	zoneDateTime := now.In(zoneID)
	formatted := zoneDateTime.String()

	nc.Publish("time.us.east", []byte(formatted))
	nc.Publish("time.us.east.atlanta", []byte(formatted))

	zoneID, err = time.LoadLocation("Europe/Warsaw")
	if err != nil {
		log.Fatal(err)
	}
	zoneDateTime = now.In(zoneID)
	formatted = zoneDateTime.String()

	nc.Publish("time.eu.east", []byte(formatted))
	nc.Publish("time.eu.east.warsaw", []byte(formatted))

	// [end wildcard_tester]
}
