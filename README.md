## pipecrypt: the SNI pass-through proxy

``pipecrypt`` is a simple TCP proxy that filters and forwards encrypted traffic
following the TLS SNI handshake, replaying the ClientHello and establishing a secure channel between
two endpoints which may be unable or unwilling to otherwise establish direct communications with each other.

This module is a intended as a testbed for research around decentralization and
inspired by the transport layer in the [Fabio](https://github.com/fabiolb/fabio) load
balancing HTTPS and TCP router.

## Synopsis

<img src="https://netblocks.org/files/netblocks-logo.png" width="200px" align="right" />

There are compelling privacy benefits associated with self-hosting services and
storage of personal data. However, exposing home networks IP addresses can introduce risk.

SNI-brokered TLS forwarding offers a simple means to maintain end-to-end security, establishing a direct
HTTPS connection between the user agent and home server, via an intermediate server
that has no means to decrypt or access the content of the communication.

This package is maintained as part of the the
[NetBlocks.org](https://netblocks.org) network observation framework.


## Installation

```sh
go build
```

