#!/usr/bin/env bash

set -e

test $(go vet ./... 2>&1 | wc -l) -gt 0                                                                                                                              && echo "go vet    failed" && exit 1 || echo "go vet    succeeded"
#test $(${GOPATH}/bin/golint ./... 2>&1 | wc -l) -gt 0                                                                                                                && echo "golint    failed" && exit 1 || echo "golint    succeeded"
test $(find . -type f -name '*'    -not -path "*.pb.go" -not -path "./.git/*" -not -path "./.workspace/*" | xargs ${GOPATH}/bin/misspell 2>&1 | wc -l) -gt 0         && echo "misspell  failed" && exit 1 || echo "misspell  succeeded"
test $(find . -type f -name '*.go' -not -path "*.pb.go" -not -path "./.git/*" -not -path "./.workspace/*" | xargs ${GOPATH}/bin/gocyclo -over 16 2>&1 | wc -l) -gt 0 && echo "gocyclo   failed" && exit 1 || echo "gocyclo   succeeded"
test $(find . -type f -name '*.go' -not -path "*.pb.go" -not -path "./.git/*" -not -path "./.workspace/*" | xargs gofmt -l -s 2>&1 | wc -l) -gt 0                    && echo "gofmt     failed" && exit 1 || echo "gofmt     succeeded"
