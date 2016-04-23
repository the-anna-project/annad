.PHONY: all anna annactl goclean gofmt goget gotest

GOPATH := ${PWD}/.workspace
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
	@# Install project dependencies.
	@mkdir -p ${PWD}/.workspace/src/github.com/xh3b4sd/
	@ln -fs ${PWD} ${PWD}/.workspace/src/github.com/xh3b4sd/
	@go get -d -v ./...
	@# Install dev dependencies.
	@go get github.com/client9/misspell/cmd/misspell
	@go get github.com/fzipp/gocyclo
	@go get github.com/golang/lint/golint

gotest:
	# Run unit tests.
	@./go.test.sh
	@echo -n "\n"
	# Run project checks.
	@GOPATH=${GOPATH} ./project.check.sh
