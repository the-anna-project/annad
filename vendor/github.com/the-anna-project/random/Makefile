.PHONY: all clean gofmt goget gotest setup



export SHELL := /bin/bash
export GOPATH := $(PWD)/.workspace



all:
	@go build \
		github.com/the-anna-project/random

clean:
	@rm -rf coverage.txt profile.out .workspace/

gofmt:
	@go fmt ./...

goget:
	@# Setup workspace.
	@mkdir -p $(PWD)/.workspace/src/github.com/the-anna-project/
	@ln -fs $(PWD) $(PWD)/.workspace/src/github.com/the-anna-project/
	@# Fetch dev dependencies.
	@#go get -u github.com/org/repo
	@# Fetch the rest of the project dependencies.
	@go get -d -v ./...

gotest:
	@./gotest

setup: goget
