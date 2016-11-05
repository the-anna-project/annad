# makefile
> A [makefile](https://en.wikipedia.org/wiki/Makefile) is a file containing a
> set of directives used with the make build automation tool.

The makefile is used to ease the development of the project. You may want to
make use of the following targets.

### all
This is the best to start with. The target fetches dependencies like `goget`
and compiles the binaries like `annad` and `annactl`.
```
make all
```

The `all` target is the default target. So the following is shorter and also
valid.
```
make
```

### annad
This target compiles the server binary `annad`. This binary can be executed to
launch Anna as process running on your system.
```
make annad
```

### annactl
This target compiles the client binary `annactl`. This binary can be executed
to interact with Anna over network.
```
make annactl
```

### dockerimage
This target builds a docker image using the [Dockerfile](/Dockerfile).
Therefore docker as dev dependency is required.
```
make dockerimage
```

### dockerpush
This target pushes a docker image to the docker repository under
https://hub.docker.com/r/xh3b4sd/anna/. Therefore docker as dev dependency is
required as well as access to the docker repository. This might be only
interesting for the maintainers.
```
make dockerpush
```

### clean
This removes the `.workspace/` directory and other files maybe flying around.
```
make clean
```

### gofmt
This is for manual code formatting. There is probably no need to use this
target on a regular basis in case there already is some automated formatting
integrated into your text editor.
```
make gofmt
```

### goget
This fetches all dependencies required for the projects source code and the
development process.
```
make goget
```

Now what happened? Because of our directory structure and `Makefile` magic the
`.workspace` directory was created and extended with the source code
dependencies.
```
~/projects/private/anna
├── ...
└── .workspace
    ├── ...
    └── src
        ├── github.com
        │   ├── ...
        ├── golang.org
        │   └── ...
        └── gopkg.in
            └── ...
```

Further some development tools were installed, necessary for some other
Makefile targets.
```
~/projects/private/anna
├── ...
└── .workspace
    ├── ...
    └── bin
        ├── ...
        ├── gocyclo
        ├── golint
        └── misspell
```

### gotest
This runs the golang unit tests to ensure basic functionality and a decent coverage.
```
make gotest
```

You should see something like this.
```
ok    _/home/vagrant/projects/private/anna/client/control/log  0.020s  coverage: 100.0% of statements
ok    _/home/vagrant/projects/private/anna/client/interface/text  0.005s  coverage: 30.0% of statements
ok    _/home/vagrant/projects/private/anna/connection-path  0.005s  coverage: 97.3% of statements
ok    _/home/vagrant/projects/private/anna/factory/id  0.040s  coverage: 91.2% of statements
ok    _/home/vagrant/projects/private/anna/factory/permutation  0.005s  coverage: 100.0% of statements
ok    _/home/vagrant/projects/private/anna/factory/random  0.027s  coverage: 100.0% of statements
ok    _/home/vagrant/projects/private/anna/index/clg/collection  2.566s  coverage: 94.6% of statements
ok    _/home/vagrant/projects/private/anna/index/clg/collection/distribution  0.009s  coverage: 85.7% of statements
ok    _/home/vagrant/projects/private/anna/index/clg/collection/feature-set  0.411s  coverage: 96.7% of statements
ok    _/home/vagrant/projects/private/anna/log  0.007s  coverage: 99.0% of statements
ok    _/home/vagrant/projects/private/anna/storage/memory  0.013s  coverage: 97.9% of statements
ok    _/home/vagrant/projects/private/anna/storage/redis  0.014s  coverage: 97.9% of statements
```

### projectcheck
This is to ensure some QA aspects like spelling, source code formatting and
others. Check `project.check.sh` for the details. It is recommended to make use
of this during your development cycles. Checks done here are also done during
CI builds. Using the following command every now and then prevents pushing
faulty branches.
```
make projectcheck
```

You should see something like this.
```
go vet    succeeded
golint    succeeded
misspell  succeeded
gocyclo   succeeded
gofmt     succeeded
```
