# development
This document describes the development environment and its processes.

### setup
At first, clone the repository.
```
git clone git@github.com:xh3b4sd/anna.git
```

To setup the go environment just execute the following in the project root
directory. This only sets some environment variables needed for the go tools.
```
. ./gosetup
```

Just running the code without installing the binary can be achieved using the
go tools.
```
go run ./src/github.com/xh3b4sd/anna/main.go
```

Installing the binary, if not yet done, can be achieved using the go tools.
```
go install ./src/github.com/xh3b4sd/anna/main.go
```
