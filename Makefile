.PHONY: all client gobuild goclean gofmt goget gorun gotest server

GOPATH := ${PWD}/.workspace/:${GOPATH}
export GOPATH

all: goget client server

client:
	@go build -o .workspace/bin/annactl github.com/xh3b4sd/anna/annactl

goclean:
	@rm -rf .workspace/

gofmt:
	@go fmt ./...

goget:
	@mkdir -p .workspace/src/github.com/xh3b4sd/
	@ln -fs ${PWD} ${PWD}/.workspace/src/github.com/xh3b4sd/
	@go get -v github.com/xh3b4sd/anna
	@go get -v github.com/xh3b4sd/anna/annactl

gotest:
	@go test ./...

server:
	@go build -o .workspace/bin/anna github.com/xh3b4sd/anna
