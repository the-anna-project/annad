.PHONY: all gobuild goclean gofmt goget gorun gotest

GOPATH := ${PWD}/.workspace/:${GOPATH}
export GOPATH

all: goclean goget

gobuild:
	@go build -o .workspace/bin/anna github.com/xh3b4sd/anna

goclean:
	@rm -rf .workspace/

gofmt:
	@go fmt ./...

goget:
	@mkdir -p .workspace/src/github.com/xh3b4sd/
	@ln -s ${PWD} ${PWD}/.workspace/src/github.com/xh3b4sd/
	@go get -v github.com/xh3b4sd/anna

gorun:
	@go run main.go

gotest:
	@go test ./...
