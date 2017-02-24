# Overview
[Competing Consumers](http://www.enterpriseintegrationpatterns.com/patterns/messaging/CompetingConsumers.html)
	
	Competing Consumers are multiple consumers that are all created to receive messages from a single Point-to-Point Channel. When the channel delivers a message, any of the consumers could potentially receive it. The messaging system's implementation determines which consumer actually receives the message, but in effect the consumers compete with each other to be the receiver. Once a consumer receives a message, it can delegate to the rest of its application to help process the message. (This solution only works with Point-to-Point Channels; multiple consumers on a Publish-Subscribe Channel just create more copies of each message.)


# Steps

1. Run a gnatsd server. `go get github.com/gnatsd; gnatsd &`
1. Get the example NATS queue subscriber. `go get github.com/nats-io/go-nats-examples/nats-qsub; nats-qsub channel_name queue_group &; nats-qsub channel_name queue_group`
1. Publish a message to the queue group. `go get github.com/nats-io/go-nats-examples/nats-pub; nats-pub channel_name message1; nats-pub channel_name message2;`

# Further Reading 


