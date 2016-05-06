.PHONY: all anna annactl goclean gofmt gogenerate goget gotest projectcheck

GOPATH := ${PWD}/.workspace
export GOPATH

all: goget annactl anna

anna: gogenerate
	@go build \
		-o .workspace/bin/anna \
		-ldflags "-X main.version=$(shell git rev-parse --short HEAD)" \
		github.com/xh3b4sd/anna

annactl: gogenerate
	@go build \
		-o .workspace/bin/annactl \
		-ldflags "-X main.version=$(shell git rev-parse --short HEAD)" \
		github.com/xh3b4sd/anna/annactl

goclean:
	@rm -rf coverage.txt profile.out .workspace/

gofmt:
	@go fmt ./...

gogenerate:
	@go generate ./...

goget:
	@# Install project dependencies.
	@mkdir -p ${PWD}/.workspace/src/github.com/xh3b4sd/
	@ln -fs ${PWD} ${PWD}/.workspace/src/github.com/xh3b4sd/
	@go get -d -v ./...
	@go get github.com/xh3b4sd/loader
	@# Install dev dependencies.
	@go get github.com/client9/misspell/cmd/misspell
	@go get github.com/fzipp/gocyclo
	@go get github.com/golang/lint/golint

gotest: gogenerate
	@./go.test.sh \

projectcheck:
	@./project.check.sh
