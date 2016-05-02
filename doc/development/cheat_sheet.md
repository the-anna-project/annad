# cheat sheet
Here are we collect some useful snippets to ease some pains that may occur
during development.

### run single tests
The following snippet can be used to easily run unit tests of single packages.
This is unlike to the makefile target, that runs the tests of all packages.
Here all tests of the package `pkg` will be run. For more options consider
reading `go test -h`.
```
GOPATH=$(pwd)/.workspace/ go test ./pkg
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
