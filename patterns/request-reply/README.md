# Overview
[Request-Reply](http://www.enterpriseintegrationpatterns.com/patterns/messaging/RequestReply.html)

When an application sends a request and expects a response from a receiver, create two channels one for the request and one for the response.

# Request-Reply 
To demonstrate the request-reply pattern with 
[gnatsd](http://www.github.com/nats-io/gnatsd) or a
[NATS.cloud](https://www.nats.cloud) instance:

  1. Get and run nats-qrply:
     ```
     go get github.com/nats-io/go-nats-examples/nats-qrly
     nats-qrply -s tls://user:password@server:port subject group "response" 
     ```
  1. Get and run nats-req:
     ```
     go get github.com/nats-io/go-nats-examples/nats-req
     nats-req -s tls://user:password@server:port subject "request" 
     ```
  1. Verify requester output:
     ```
     Published [subject] : 'request'
     Received [_INBOX.5cE7rtqAFLWkhaLEawkBqL] : 'response'
     ```
  1. Verify queue reply output:
     ```
     Listening on [subject] [Group: group]
     [#1] Received on [subject]: 'request'
     ```

Now that you've seen how it works, try running multiple `nats-qrply` and sending multiple messages. For each request, a random queue replier will be selected and at most one response will be sent back to the client.
