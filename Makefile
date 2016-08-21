.PHONY: all anna annactl dockerimage dockerpush goclean gofmt gogenerate goget gotest projectcheck



GOPATH := ${PWD}/.workspace
export GOPATH



VERSION := $(shell git rev-parse --short HEAD)



all: annactl anna

anna: goget gogenerate
	@go build \
		-o .workspace/bin/anna \
		-ldflags "-X main.version=${VERSION}" \
		github.com/xh3b4sd/anna/anna

annactl: goget gogenerate
	@go build \
		-o .workspace/bin/annactl \
		-ldflags "-X main.version=${VERSION}" \
		github.com/xh3b4sd/anna/annactl

dockerimage: all
	@docker build -t xh3b4sd/anna:${VERSION} .

dockerpush:
	docker push xh3b4sd/anna:${VERSION}

goclean:
	@rm -rf coverage.txt profile.out .workspace/

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

setup: protoc

projectcheck:
	@./project.check.sh

protoc:
	@wget https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip
	@unzip protoc-3.0.0-linux-x86_64.zip -d protoc
	@mv protoc/bin/protoc /usr/local/bin/protoc
	@rm -rf protoc protoc-3.0.0-linux-x86_64.zip
