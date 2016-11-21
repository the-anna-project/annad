[![maturity](https://img.shields.io/badge/status-alpha-red.svg)](https://github.com/the-anna-project/client) [![build status](https://travis-ci.org/the-anna-project/client.svg?branch=master)](https://travis-ci.org/the-anna-project/client) [![GoDoc](https://godoc.org/github.com/the-anna-project/client?status.svg)](http://godoc.org/github.com/the-anna-project/client)

# client
The client service provides an endpoint collection for Anna's network API.

### build
This project uses Protocol Buffers and gRPC code generation. The `Makefile`
helps to get this right.

```
make devdeps
make gogenerate
```
