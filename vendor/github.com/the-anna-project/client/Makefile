.PHONY: all devdeps gogenerate



export SHELL := /bin/bash
export PATH := ${PATH}:${GOPATH}/bin



all: devdeps gogenerate

devdeps:
	@go get -u -v github.com/golang/protobuf/protoc-gen-go
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

gogenerate:
	@protoc --proto_path=${GOPATH}/src/github.com/the-anna-project/spec/api --go_out=plugins=grpc,import_path=text:service/text/ ${GOPATH}/src/github.com/the-anna-project/spec/api/text_endpoint.proto
