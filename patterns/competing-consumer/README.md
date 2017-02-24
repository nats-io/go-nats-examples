# Overview
[Competing Consumers](http://www.enterpriseintegrationpatterns.com/patterns/messaging/CompetingConsumers.html)
	
"Competing Consumers are multiple consumers that are all created to receive messages from a single Point-to-Point Channel. When the channel delivers a message, any of the consumers could potentially receive it. The messaging system's implementation determines which consumer actually receives the message, but in effect the consumers compete with each other to be the receiver. Once a consumer receives a message, it can delegate to the rest of its application to help process the message. (This solution only works with Point-to-Point Channels; multiple consumers on a Publish-Subscribe Channel just create more copies of each message.)"


# Steps

1. Run a gnatsd server. 
	
	```
	go get github.com/gnatsd; gnatsd &
	```
	
1. Install the queue subscriber. 
	
	```
	go get github.com/nats-io/go-nats-examples/nats-qsub
	```
	
1. Run the first queue_group subscriber. 
	
	```
	nats-qsub subject queue_group
	```
	
1. In a separate terminal window, run the second queue_group subscriber. 
	
	```
	nats-qsub subject queue_group
	```
	
1. In a 3rd terminal window, send messages to the queue group. 
	
	```
	go get github.com/nats-io/go-nats-examples/nats-pub
	nats-pub subject message1
	nats-pub subject message2
	```

# Further Reading 
