# makefile
> A [makefile(https://en.wikipedia.org/wiki/Makefile) is a file containing a
> set of directives used with the make build automation tool.

The makefile is used to ease the development of the project. You may want to
make use of the following targets.

### all
This is the best to start with. The target fetches dependencies like `goget`
and compiles the binaries like `anna` and `annactl`.
```
make all
```

The `all` target is the default target. So the following is shorter and also
valid.
```
make
```

### anna
This target compiles the server binary `anna`. This binary can be executed to
launch Anna as process running on your system.
```
make anna
```

### annactl
This target compiles the client binary `annactl`. This binary can be executed
to interact with Anna over network.
```
make annactl
```

### goclean
This removes dependencies to cleanup the `.workspace/` directory.
```
make goclean
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
ok    _/home/vagrant/projects/private/anna/clg  1.955s  coverage: 98.4% of statements
ok    _/home/vagrant/projects/private/anna/clg/distribution 0.008s  coverage: 100.0% of statements
ok    _/home/vagrant/projects/private/anna/clg/feature-set  0.388s  coverage: 99.3% of statements
ok    _/home/vagrant/projects/private/anna/client/control/log 0.026s  coverage: 96.0% of statements
ok    _/home/vagrant/projects/private/anna/client/interface/text  0.020s  coverage: 98.1% of statements
ok    _/home/vagrant/projects/private/anna/factory/client 0.006s  coverage: 100.0% of statements
ok    _/home/vagrant/projects/private/anna/factory/server 0.010s  coverage: 96.1% of statements
ok    _/home/vagrant/projects/private/anna/gateway  0.008s  coverage: 100.0% of statements
ok    _/home/vagrant/projects/private/anna/id 0.043s  coverage: 87.5% of statements
ok    _/home/vagrant/projects/private/anna/log  0.005s  coverage: 99.0% of statements
ok    _/home/vagrant/projects/private/anna/net/pat  0.006s  coverage: 90.0% of statements
ok    _/home/vagrant/projects/private/anna/scheduler  0.178s  coverage: 92.1% of statements
ok    _/home/vagrant/projects/private/anna/storage/memory 0.011s  coverage: 98.0% of statements
ok    _/home/vagrant/projects/private/anna/storage/redis  0.013s  coverage: 96.8% of statements
ok    _/home/vagrant/projects/private/anna/strategy 0.007s  coverage: 98.5% of statements
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
