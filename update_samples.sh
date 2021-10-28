#!/bin/bash

nats_tools=(nats-bench nats-echo nats-pub nats-qsub nats-req nats-rply nats-sub)
for i in "${nats_tools[@]}"
do
    echo "downloading ${i}"
	curl "https://raw.githubusercontent.com/nats-io/nats.go/main/examples/${i}/main.go" -o "tools/${i}/${i}.go"
done

stan_tools=(stan-pub stan-sub)
for i in "${stan_tools[@]}"
do
    echo "downloading ${i}"
	curl "https://raw.githubusercontent.com/nats-io/stan.go/main/examples/${i}/main.go" -o "tools/${i}/${i}.go"
done
