.PHONY: all



GOPATH := ${PWD}/.workspace
export GOPATH



all: goget lib

clean:
	rm -rf ${PWD}/.workspace worker-pool

goget:
	@# setup workspace.
	@mkdir -p ${PWD}/.workspace/src/github.com/xh3b4sd/
	@ln -fs ${PWD} ${PWD}/.workspace/src/github.com/xh3b4sd/
	@# install project dependencies.
	@go get -d -v ./...

lib:
	@go build
