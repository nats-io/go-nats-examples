![](https://raw.githubusercontent.com/nats-io/nats-site/master/src/img/large-logo.png)

# NATS - Go Examples and CLI Clients

[Go](http://www.golang.org) examples and CLI clients for the [NATS messaging system](https://nats.io).

[![License Apache 2](https://img.shields.io/badge/License-Apache2-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Build Status](https://travis-ci.org/nats-io/go-nats-examples.svg?branch=master)](http://travis-ci.org/nats-io/go-nats-examples)

# Overview
This repo contains go-gettable [go-nats](www.github.com/nats-io/go-nats) example client code as well as example code from the documentation.

# Verify a NATS Server
To verify a NATS Server with either your own
[gnatsd](www.github.com/nats-io/gnatsd) or a
[NATS.cloud](www.nats.cloud) instance:

  1. Get and run nats-sub:
     ```
     go get github.com/nats-io/go-nats-examples/tools/nats-sub
     nats-sub -s tls://user:password@server:port channel_name
     ```
  1. Get and run nats-pub:
     ```
     go get github.com/nats-io/go-nats-examples/tools/nats-pub
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

# Patterns
The patterns directory contains a listing of example messaging patterns:

  1. [Publish/Subscribe](/patterns/publish-subscribe)
  1. [Request/Reply](/patterns/request-reply)
  1. [Competing Consumer](/patterns/competing-consumer/)
