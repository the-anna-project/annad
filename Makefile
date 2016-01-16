.PHONY: gobuild goclean gofmt goget gorun gotest

PROJECT := github.com/xh3b4sd/anna

GOPATH := ${PWD}/vendor:${GOPATH}
export GOPATH

gobuild:
	@go build -o ./bin/anna ./src/${PROJECT}

goclean:
	@rm -rf ./bin
	@rm -rf ./pkg
	@rm -rf ./vendor

gofmt:
	@go fmt ./src/${PROJECT}/...

goget:
	@mkdir -p ./vendor/
	@go get -v ./src/${PROJECT}/...

gorun:
	@go run ./src/${PROJECT}/main.go

gotest:
	@go test ./src/${PROJECT}/...
