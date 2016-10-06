# server
Documentation about the server described here refers to the [server
package](https://godoc.org/github.com/xh3b4sd/anna/server) written in
[golang](https://golang.org). The server provides Anna's API to execute her
business logic from remote. The [client](client.md) may communicate with the
server to ask for any kind of operation of Anna's [interface](interface.md) or
or to [control](control.md) certain administrative behaviour. The current
implementation of the client and the server makes use of
https://github.com/go-kit/kit for the HTTP endpoints and
https://github.com/grpc/grpc-go for the gRPC endpoints.
