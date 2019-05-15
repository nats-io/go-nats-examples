#!/bin/bash

nats_tools=(nats-bench nats-echo nats-pub nats-qsub nats-req nats-rply nats-sub)
for i in "${nats_tools[@]}"
do
	curl "https://raw.githubusercontent.com/nats-io/nats.go/master/examples/${i}/main.go" -o "tools/${i}/${i}.go"
done

stan_tools=(stan-pub stan-sub)
for i in "${stan_tools[@]}"
do
	curl "https://raw.githubusercontent.com/nats-io/go-nats-streaming/master/examples/${i}/main.go" -o "tools/${i}/${i}.go"
done