.PHONY: all annad annactl clean dockerimage dockerpush gofmt gogenerate goget gotest projectcheck protoc setup



export SHELL := /bin/bash
export GOPATH := $(PWD)/.workspace
export PATH := $(PATH):$(PWD)/.workspace/bin:$(PWD)/bin


VERSION := $(shell git rev-parse --short HEAD)



all: annactl annad

annad: gogenerate
	@go build \
		-o .workspace/bin/annad \
		-ldflags "-X main.version=${VERSION}" \
		github.com/xh3b4sd/anna/annad

annactl: gogenerate
	@go build \
		-o .workspace/bin/annactl \
		-ldflags "-X main.version=${VERSION}" \
		github.com/xh3b4sd/anna/annactl

clean:
	@rm -rf coverage.txt profile.out .workspace/ /tmp/protoc/ /tmp/protoc.zip
	@# TODO remove generated code

dockerimage: all
	@docker build -t xh3b4sd/anna:${VERSION} .

dockerpush:
	docker push xh3b4sd/anna:${VERSION}

gofmt:
	@go fmt ./...

gogenerate:
	@go generate ./...
	@protoc --proto_path=spec --go_out=plugins=grpc,import_path=text:client/interface/text/ spec/textendpoint.proto
	@protoc --proto_path=spec --go_out=plugins=grpc,import_path=textendpoint:service/textendpoint/ spec/textendpoint.proto

goget:
	@# Setup workspace.
	@mkdir -p $(PWD)/.workspace/src/github.com/xh3b4sd/
	@ln -fs $(PWD) $(PWD)/.workspace/src/github.com/xh3b4sd/
	@# Fetch dev dependencies.
	@go get -u github.com/client9/misspell/cmd/misspell
	@go get -u github.com/fzipp/gocyclo
	@go get -u github.com/golang/lint/golint
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@go get -u github.com/xh3b4sd/clggen
	@# Fetch the rest of the project dependencies.
	@go get -d -v ./...

gotest: gogenerate
	@gotest

projectcheck:
	@projectcheck

protoc:
ifeq ($(shell go env GOOS),linux)
	@wget https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip -O /tmp/protoc.zip
else ifeq ($(shell go env GOOS),darwin)
	@wget https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-osx-x86_64.zip -O /tmp/protoc.zip
else
	@echo "unsupported platform"
	@exit 1
endif
	@unzip /tmp/protoc.zip -d /tmp/protoc/
	@mkdir -p .workspace/bin/
	@mv /tmp/protoc/bin/protoc .workspace/bin/protoc
	@rm -rf /tmp/protoc/ /tmp/protoc.zip

setup: goget protoc
