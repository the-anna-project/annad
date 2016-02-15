.PHONY: all client gobuild goclean gofmt goget gorun gotest server

$(mkdir -p .workspace/)
GOPATH := ${PWD}/.workspace/:${PWD}/../../../..:${GOPATH}
export GOPATH

all: goget client server

client:
	@go build -o .workspace/bin/annactl github.com/xh3b4sd/anna/annactl

goclean:
	@rm -rf .workspace/

gofmt:
	@go fmt ./...

goget:
	@go get -v github.com/xh3b4sd/anna
	@go get -v github.com/xh3b4sd/anna/annactl

gotest:
	@go test ./... -cover

server:
	@go build -o .workspace/bin/anna github.com/xh3b4sd/anna
