.PHONY: all annad clean devdeps dockerimage dockerpush gofmt gogenerate gotest projectcheck setup



export SHELL := /bin/bash



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
	@rm -rf coverage.txt profile.out
	@# TODO remove generated code

devdeps:
	@# Fetch dev dependencies.
	@go get -u -v github.com/client9/misspell/cmd/misspell
	@go get -u -v github.com/fzipp/gocyclo
	@go get -u -v github.com/golang/lint/golint
	@go get -u -v github.com/xh3b4sd/clggen

dockerimage: all
	@docker build -t xh3b4sd/anna:${GIT_COMMIT} .

dockerpush:
	docker push xh3b4sd/anna:${GIT_COMMIT}

gofmt:
	@go fmt ./...

gogenerate:
	@go generate ./service/clg

gotest: gogenerate
	@./bin/gotest

projectcheck:
	@./bin/projectcheck

setup: devdeps
