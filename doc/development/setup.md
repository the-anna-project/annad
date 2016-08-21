# setup
In https://golang.org the project structure is something you need to deal with.
When you ask 5 people how to do it, you probably get six answers. Here I simply
describe my personal workflow and how the project is set up using the
[makefile](makefile.md). Setup the project as you like so it fits your
workflow. Anyway, this is how I am doing it.

In golang the `GOPATH` assumes that there is something like
`src/github.com/xh3b4sd/anna/` within your workspace. See
https://golang.org/doc/code.html. I am using the [makefile](makefile.md) to
execute all go commands. Note how the `GOPATH` is set in the makefile.
```
GOPATH := ${PWD}/.workspace/
export GOPATH
```

This takes care that the workspace is properly set up for the dependencies that
`go get` fetches and that `GOPATH` itself is properly set. So the result after
the setup will be similar to this.
```
~/projects/private/anna    # holds the project's source code
├── ...
└── .workspace             # represents the GOPATH holding the project's dependencies
    └── ...
```

### install protocol-buffers

In order to install `protoc`, simply download a prebuild binary from here:
https://github.com/google/protobuf/releases.
```
wget https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip
unzip protoc-3.0.0-linux-x86_64.zip -d protoc
mv protoc/bin/protoc /usr/local/bin/protoc
rm -rf protoc protoc-3.0.0-linux-x86_64.zip
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

### clone repository
To setup the project I am creating the working directory and go into it. It
holds the source code of the Anna project.
```
mkdir -p ~/projects/private/anna/
cd ~/projects/private/anna/
```

Then, I clone the repository. Note the `.` at the end of the command.
```
git clone git@github.com:xh3b4sd/anna.git .
```

### build
I am building the project using the [makefile](makefile.md). See its
documentation for more information on how to use it.
```
make
```
