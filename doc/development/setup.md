# setup
In https://golang.org the project structure is something you need to deal with.
See also https://golang.org/doc/code.html. So it is best to have one single
`$GOPATH` in which all your projects live in. To more easily manage a complex
project like this one here we are using the [makefile](makefile.md) to execute
all go commands. In my `.zshrc` I set the `$GOPATH` like this.
```
export GOPATH=~/gopath
```

The Anna project is then located here.
```
~/gopath/src/github.com/the-anna-project/annad    
```

### prerequisites
The Anna project requires
[protocol-buffers](https://developers.google.com/protocol-buffers/). In order
to install `protoc`, simply execute the following command.
```
make setup
```

When executing the `protoc` library you might experience the following error.
```
protoc: error while loading shared libraries: libprotoc.so.10: cannot open shared object file: No such file or directory
```

In such a case the following command fixed this issue for me. See also
http://stackoverflow.com/questions/25518701/protobuf-cannot-find-shared-libraries.
```
sudo ldconfig
```

### build
I am building the project using the [makefile](makefile.md). See its
documentation for more information on how to use it.
```
make all
```
