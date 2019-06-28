![](https://raw.githubusercontent.com/nats-io/nats-site/master/src/img/large-logo.png)

# NATS - Go Examples and CLI Clients

[Go](http://www.golang.org) examples and CLI clients for the [NATS messaging system](https://nats.io).

[![License Apache 2](https://img.shields.io/badge/License-Apache2-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Build Status](https://travis-ci.org/nats-io/go-nats-examples.svg?branch=master)](http://travis-ci.org/nats-io/go-nats-examples)

# Overview
This repo contains go-gettable [nats.go](www.github.com/nats-io/nats.go) examples and client code as well as api examples from the documentation.

Install your own server, or optionally utilize the [demo server](http://demo.nats.io:8222)

  1. Get and run nats-sub:
     ```
     go get github.com/nats-io/go-nats-examples/tools/nats-sub
     nats-sub -s demo.nats.io subject_name
     ```
  1. Get and run nats-pub:
     ```
     go get github.com/nats-io/go-nats-examples/tools/nats-pub
     nats-pub -s demo.nats.io subject_name msg
     ```
  1. Verify publisher output:
     ```
     Published [subject_name] : 'message'
     ```
  1. Verify subscriber output:
     ```
     Listening on [subject_name]
     [#1] Received on [subject_name]: 'message'
     ```

# Patterns
The patterns directory contains a listing of example messaging patterns:

  1. [Publish/Subscribe](/patterns/publish-subscribe)
  1. [Request/Reply](/patterns/request-reply)
  1. [Competing Consumer](/patterns/competing-consumer/)
