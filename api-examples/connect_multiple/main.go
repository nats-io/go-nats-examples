package main

import (
	"log"
	"strings"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_multiple]
	servers := []string{"nats://localhost:1222",
		"nats://localhost:1223",
		"nats://localhost:1224",
	}

	nc, err := nats.Connect(strings.Join(servers, ","))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_multiple]
}
