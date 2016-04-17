.PHONY: all anna annactl goclean gofmt goget gotest

GOPATH := ${PWD}/.workspace/
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

cicover:
	go get github.com/mattn/goveralls
	go get golang.org/x/tools/cmd/cover
	.workspace/bin/goveralls -service=travis-ci

goclean:
	@rm -rf .workspace/ coverage.txt

gofmt:
	@go fmt ./...

goget:
	@mkdir -p ${PWD}/.workspace/src/github.com/xh3b4sd/
	@ln -fs ${PWD} ${PWD}/.workspace/src/github.com/xh3b4sd/
	@go get -d -v ./...

gotest:
	@./go.test.sh
	@go vet ./...
