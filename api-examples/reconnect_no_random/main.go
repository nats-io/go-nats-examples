package main

import (
	"log"
	"strings"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin reconnect_no_random]
	servers := []string{"nats://127.0.0.1:1222",
		"nats://127.0.0.1:1223",
		"nats://127.0.0.1:1224",
	}

	nc, err := nats.Connect(strings.Join(servers, ","), nats.DontRandomize())
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end reconnect_no_random]
}
