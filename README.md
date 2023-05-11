# Fork of QTLS-GO1-20

This repository contains a modified version of [qtls-go1-19](https://github.com/quic-go/qtls-go1-20).

## Security
This implementation is intended for research purposes only and should not be deployed on the internet.

## Changes to the original QTLS-GO1-19
- FromTrafficSecret creates a new TLS connection without doing a handshake
- Sent and received records can be observed