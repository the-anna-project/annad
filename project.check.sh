#!/usr/bin/env bash

set -e

test $(go vet ./... 2>&1 | wc -l) -gt 0                                                                                   && echo "go vet    failed" && exit 1 || echo "go vet    succeeded"
test $(golint ./... 2>&1 | wc -l) -gt 0                                                                                   && echo "golint    failed" && exit 1 || echo "golint    succeeded"
test $(find . -name '*'    -not -path "./.git/*" -not -path "./.workspace/*" | xargs misspell 2>&1 | wc -l) -gt 0         && echo "misspell  failed" && exit 1 || echo "misspell  succeeded"
test $(find . -name '*.go' -not -path "./.git/*" -not -path "./.workspace/*" | xargs gocyclo -over 15 2>&1 | wc -l) -gt 0 && echo "gocyclo   failed" && exit 1 || echo "gocyclo   succeeded"
test $(find . -name '*.go' -not -path "./.git/*" -not -path "./.workspace/*" | xargs gofmt -l -s 2>&1 | wc -l) -gt 0      && echo "gofmt     failed" && exit 1 || echo "gofmt     succeeded"
