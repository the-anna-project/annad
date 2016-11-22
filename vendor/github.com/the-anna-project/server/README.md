[![maturity](https://img.shields.io/badge/status-alpha-red.svg)](https://github.com/the-anna-project/server) [![build status](https://travis-ci.org/the-anna-project/server.svg?branch=master)](https://travis-ci.org/the-anna-project/server) [![GoDoc](https://godoc.org/github.com/the-anna-project/server?status.svg)](http://godoc.org/github.com/the-anna-project/server)

# server
The server service provides an endpoint collection for Anna's network API.

### build
This project uses Protocol Buffers and gRPC code generation. The `Makefile`
helps to get this right.

```
make devdeps
make gogenerate
```
