# cheat sheet
Here are we collect some useful snippets to ease some pains that may occur
during development.

### fix misspells
There may be the project check failing like follows.
```
$ make projectcheck
go vet    succeeded
golint    succeeded
misspell  failed
make: *** [projectcheck] Error 1
```

The following snippet can be used to easily check misspells as reported by
https://github.com/client9/misspell.
```
find . -not -path "./.git/*" -not -path "./.workspace/*" | xargs misspell
```

You will see output similar to the following depending on the current issues.
```
file:388:4:found "wrong." a misspelling of "right."
file:404:21:found "wrong" a misspelling of "right"
file:4:found "wrong." a misspelling of "right."
file:21:found "wrong" a misspelling of "right"
```

The following snippet can be used to easily correct misspells. Simply add `-w`
to the command we used to check the misspells. The output will be the same as
already seen and the corrections will be written to the related files.
```
find . -not -path "./.git/*" -not -path "./.workspace/*" | xargs misspell -w
```

### run single tests
The following snippet can be used to easily run unit tests of single packages.
This is unlike to the makefile target, that runs the tests of all packages.
Here all tests of the package `pkg` will be run. For more options consider
reading `go test -h`. Note that it is a good idea to run tests with having the
race detector enabled. For more information on this one see
https://blog.golang.org/race-detector.
```
GOPATH=$(pwd)/.workspace/ go test -race ./pkg
```

### search and replace
The following snippet can be used to easily search and replace strings
recursively within a given directory. Here `search` is replaced by `replace`
within all files which names match the expression `*.go` within the current
directory given by `.`. To accomplish that we make use of the command line
tools `sed` and `find`.
```
sed -i 's/search/replace/g' $(find . -name *.go)
```

### list package dependencies
Sometimes it can be helpful to know which packages are imported within a
specific go package. `go list` is your friend. Note that there are also other
interesting information within the provided JSON output besides `Imports`. Try
e.g. `Deps`.
```
$ go list -json ./service/random | jq .Imports
[
  "crypto/rand",
  "fmt",
  "github.com/cenk/backoff",
  "github.com/juju/errgo",
  "github.com/the-anna-project/annad/service/spec",
  "io",
  "math/big",
  "time"
]
```
