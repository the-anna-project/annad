.PHONY: all anna annactl clean dockerimage dockerpush gofmt gogenerate goget gotest projectcheck



GOPATH := ${PWD}/.workspace
PATH := ${PATH}:${PWD}/.workspace/bin
export GOPATH
export PATH



VERSION := $(shell git rev-parse --short HEAD)



all: annactl anna

anna: gogenerate
	@go build \
		-o .workspace/bin/anna \
		-ldflags "-X main.version=${VERSION}" \
		github.com/xh3b4sd/anna/anna

annactl: gogenerate
	@go build \
		-o .workspace/bin/annactl \
		-ldflags "-X main.version=${VERSION}" \
		github.com/xh3b4sd/anna/annactl

clean:
	@rm -rf coverage.txt profile.out .workspace/

dockerimage: all
	@docker build -t xh3b4sd/anna:${VERSION} .

dockerpush:
	docker push xh3b4sd/anna:${VERSION}

gofmt:
	@go fmt ./...

gogenerate:
	@go generate ./...
	@protoc api/text_interface.proto --go_out=plugins=grpc:.

goget:
	@# Setup workspace.
	@mkdir -p ${PWD}/.workspace/src/github.com/xh3b4sd/
	@ln -fs ${PWD} ${PWD}/.workspace/src/github.com/xh3b4sd/
	@# Install project dependencies.
	@go get -d -v ./...
	@go get github.com/xh3b4sd/clggen
	@# Install dev dependencies.
	@go get github.com/client9/misspell/cmd/misspell
	@go get github.com/fzipp/gocyclo
	@go get github.com/golang/lint/golint
	@go get github.com/golang/protobuf/proto
	@go get github.com/golang/protobuf/protoc-gen-go

gotest: gogenerate
	@./go.test.sh

setup: goget protoc

projectcheck:
	@./project.check.sh

protoc:
	@wget https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip -O /tmp/protoc.zip
	@unzip /tmp/protoc.zip -d /tmp/protoc/
	@mv /tmp/protoc/bin/protoc .workspace/bin/protoc
	@rm -rf /tmp/protoc/ /tmp/protoc.zip
