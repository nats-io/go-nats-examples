# Overview
[Request-Reply](http://www.enterpriseintegrationpatterns.com/patterns/messaging/RequestReply.html)

When an application sends a request and expects a response from a receiver, create two channels one for the request and one for the response.

# Request-Reply
To demonstrate NATS request/reply.

  1. Get and run nats-rply:
     ```
     go get github.com/nats-io/go-nats-examples/tools/nats-rply
     nats-rply -s demo.nats.io subject "my response"
     ```
  1. Get and run nats-req:
     ```
     go get github.com/nats-io/go-nats-examples/tools/nats-req
     nats-req -s demo.nats.io subject "my request"
     ```
  1. Verify requester output:
     ```
     Published [subject] : 'my request'
     Received [_INBOX.5cE7rtqAFLWkhaLEawkBqL] : 'my response'
     ```
  1. Verify queue reply output:
     ```
     Listening on [subject]
     [#1] Received on [subject]: 'my request'
     ```

Now that you've seen how it works, try running multiple `nats-rply` and sending multiple messages. For each request, a random queue responder will be selected and at most one response will be sent back to the client.
