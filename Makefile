.PHONY: all anna annactl dockerimage dockerpush goclean gofmt gogenerate goget gotest projectcheck



GOPATH := ${PWD}/.workspace
export GOPATH



VERSION := $(shell git rev-parse --short HEAD)



all: goget annactl anna

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

goget:
	@# Setup workspace.
	@mkdir -p ${PWD}/.workspace/src/github.com/xh3b4sd/
	@ln -fs ${PWD} ${PWD}/.workspace/src/github.com/xh3b4sd/
	@# Install project dependencies.
	@go get -d -v ./...
	@go get -u github.com/xh3b4sd/clggen
	@# Install dev dependencies.
	@go get -u github.com/client9/misspell/cmd/misspell
	@go get -u github.com/fzipp/gocyclo
	@go get -u github.com/golang/lint/golint

gotest: gogenerate
	@./go.test.sh

projectcheck:
	@./project.check.sh
