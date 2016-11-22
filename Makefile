.PHONY: all annad clean devdeps dockerimage dockerpush gofmt gogenerate gotest projectcheck protoc setup



export SHELL := /bin/bash
export PATH := ${PATH}:${GOPATH}/bin



GIT_COMMIT := $(shell git rev-parse --short HEAD)
ifndef GOARCH
	GOARCH := $(shell go env GOARCH)
endif
ifndef GOOS
	GOOS := $(shell go env GOOS)
endif
GO_VERSION=$(shell go version | cut -d ' ' -f 3)
PROJECT_VERSION=$(shell cat VERSION)



all: annad

annad: gogenerate
	@go build \
		-o ${GOPATH}/bin/annad \
		-ldflags " \
			-X main.gitCommit=${GIT_COMMIT} \
			-X main.goArch=${GOARCH} \
			-X main.goOS=${GOOS} \
			-X main.goVersion=${GO_VERSION} \
			-X main.projectVersion=${PROJECT_VERSION} \
		" \
		.

clean:
	@rm -rf coverage.txt profile.out /tmp/protoc/ /tmp/protoc.zip
	@# TODO remove generated code

devdeps:
	@# Fetch dev dependencies.
	@go get -u -v github.com/client9/misspell/cmd/misspell
	@go get -u -v github.com/fzipp/gocyclo
	@go get -u -v github.com/golang/lint/golint
	@go get -u -v github.com/golang/protobuf/protoc-gen-go
	@go get -u -v github.com/xh3b4sd/clggen

dockerimage: all
	@docker build -t xh3b4sd/anna:${GIT_COMMIT} .

dockerpush:
	docker push xh3b4sd/anna:${GIT_COMMIT}

gofmt:
	@go fmt ./...

gogenerate:
	@go generate ./service/clg
	@protoc --proto_path=vendor/github.com/the-anna-project/spec/api --go_out=plugins=grpc,import_path=text:service/endpoint/text/ vendor/github.com/the-anna-project/spec/api/text_endpoint.proto

gotest: gogenerate
	@./bin/gotest

projectcheck:
	@./bin/projectcheck

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
	@mv /tmp/protoc/bin/protoc ${GOPATH}/bin/protoc
	@rm -rf /tmp/protoc/ /tmp/protoc.zip

setup: devdeps protoc
