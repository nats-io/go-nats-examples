# Overview
[Publish Subscribe Channel](http://www.enterpriseintegrationpatterns.com/patterns/messaging/PublishSubscribeChannel.html)

"A Publish-Subscribe Channel works like this: It has one input channel
that splits into multiple output channels, one for each subscriber.
When an event is published into the channel, the Publish-Subscribe
Channel delivers a copy of the message to each of the output channels.
Each output channel has only one subscriber, which is only allowed to
consume a message once. In this way, each subscriber only gets the
message (at-most in the case of gnatsd) once and consumed copies disappear from their channels.
"

# Publish / Subscribe
To verify a NATS Server with either your own
[gnatsd](http://www.github.com/nats-io/gnatsd) or a
[NATS.cloud](https://www.nats.cloud) instance:

  1. Get and run nats-sub:
     ```
     go get github.com/nats-io/go-nats-examples/nats-sub
     nats-sub -s tls://user:password@server:port channel_name
     ```
  1. Get and run nats-pub:
     ```
     go get github.com/nats-io/go-nats-examples/nats-pub
     nats-pub -s tls://user:password@server:port channel_name message
     ```
  1. Verify publisher output:
     ```
     Published [channel_name] : 'message'
     ```
  1. Verify subscriber output:
     ```
     Listening on [channel_name]
     [#1] Received on [channel_name]: 'message'
     ```
