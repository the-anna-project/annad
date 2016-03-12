#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(find ./* -maxdepth 10 -type f -name '*test.go' -not -path "./.workspace/*" -not -path "./.git/*" -exec dirname {} \; | sort -u); do
  go test -coverprofile=profile.out -covermode=atomic $d
  if [ -f profile.out ]; then
      cat profile.out >> coverage.txt
      rm profile.out
  fi
done
