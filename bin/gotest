#!/usr/bin/env bash

set -e

# We truncate the whole file because we are going to put new coverage data into
# it. Note that using "echo '' > file" would cause new lines which some
# coverage uploader do not like.
> coverage.txt

for d in $(find ./* -maxdepth 10 -type f -name '*test.go' -not -path "./.workspace/*" -not -path "./.git/*" -exec dirname {} \; | sort -u); do
  go test -race -covermode=atomic -coverprofile=profile.out $d
  if [ -f profile.out ]; then
      cat profile.out >> coverage.txt
      rm profile.out
  fi
done
