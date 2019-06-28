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
To demonstrate NATS publish/subscribe.

  1. Get and run nats-sub:
     ```
     go get github.com/nats-io/go-nats-examples/tools/nats-sub
     nats-sub -s demo.nats.io some_subject
     ```
  2. Get and run nats-pub:
     ```
     go get github.com/nats-io/go-nats-examples/tools/nats-pub
     nats-pub -s demo.nats.io some_subject message
     ```
  3. Verify publisher output:
     ```
     Published [subject_name] : 'message'
     ```
  4. Verify subscriber output:
     ```
     Listening on [subject_name]
     [#1] Received on [channel_name]: 'message'
     ```

  4. Optionally install your own server:
     ```
     go get github.com/nats-io/nats-server; nats-server &
     ```
