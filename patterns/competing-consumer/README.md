# Overview
[Competing Consumers](http://www.enterpriseintegrationpatterns.com/patterns/messaging/CompetingConsumers.html)

"Competing Consumers are multiple consumers that are all created to receive messages from a single Point-to-Point Channel. When the channel delivers a message, any of the consumers could potentially receive it. The messaging system's implementation determines which consumer actually receives the message, but in effect the consumers compete with each other to be the receiver. Once a consumer receives a message, it can delegate to the rest of its application to help process the message. (This solution only works with Point-to-Point Channels; multiple consumers on a Publish-Subscribe Channel just create more copies of each message.)"


# Steps

1. Run a [nats-server](http://www.github.com/nats-io/nats-server) or connect the NATS programs to the [demo server](https://nats.demo.io:8222), e.g. ` nats-pub -s demo.nats.io subject msg`

	```
	go get github.com/nats-io/nats-server; nats-server &
	```

1. Install the queue subscriber.

	```
	go get github.com/nats-io/go-nats-examples/tools/nats-qsub
	```

2. Run the first queue_group subscriber.

	```
	nats-qsub subject group
	```

3. In a separate terminal window, run the second queue_group subscriber.

	```
	nats-qsub subject group
	```

4. In a 3rd terminal window, send messages to the queue group.

	```
	go get github.com/nats-io/go-nats-examples/tools/nats-pub
	nats-pub subject msg-1
	nats-pub subject msg-2
	```

5. The nats-qsubs will process messages one at a time.