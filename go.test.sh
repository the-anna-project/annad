#!/usr/bin/env bash

set -e

# We truncate the whole file because we are going to put new coverage data into
# it. Note that using "echo '' > file" would cause new lines which some
# coverage uploader do not like.
> coverage.txt

for d in $(find ./* -maxdepth 10 -type f -name '*test.go' -not -path "./.workspace/*" -not -path "./.git/*" -exec dirname {} \; | sort -u); do
  go test -covermode=count -coverprofile=profile.out $d
  if [ -f profile.out ]; then
      cat profile.out >> coverage.txt
      rm profile.out
  fi
done

# The coverage output of the golang coverage contains the string "mode: <mode>"
# as first statement. Coverage services uploader like the one of coveralls.io
# do not like this. Thus we remove it. Note that we need "|| true" because grep
# does not return 0 when it works as expected.
cat coverage.txt | grep -v "mode: .*" > coverage.txt || true
