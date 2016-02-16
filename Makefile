.PHONY: all anna annactl gobuild goclean gofmt goget gorun gotest

$(mkdir -p .workspace/)
GOPATH := ${PWD}/.workspace/:${PWD}/../../../..:${GOPATH}
export GOPATH

all: goget annactl anna

anna:
	@go build \
		-o .workspace/bin/anna \
		-ldflags "-X main.version=$(shell git rev-parse --short HEAD)" \
		github.com/xh3b4sd/anna

annactl:
	@go build \
		-o .workspace/bin/annactl \
		-ldflags "-X main.version=$(shell git rev-parse --short HEAD)" \
		github.com/xh3b4sd/anna/annactl

goclean:
	@rm -rf .workspace/ coverage.txt

gofmt:
	@go fmt ./...

goget:
	@go get -v github.com/xh3b4sd/anna
	@go get -v github.com/xh3b4sd/anna/annactl

gotest:
	@./go.test.sh
