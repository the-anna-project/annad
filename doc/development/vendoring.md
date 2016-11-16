# vendoring
Dependency management in go is a joke. There is actually none. There is this
idea about vendoring but there is somehow no real tooling around that provided
by the golang gang. There are loads of information about it out there. Just
google `golang vendoring`. [There are also loads of
approaches](https://github.com/golang/go/wiki/PackageManagementTools) that try
to do it. All of them are not what you probably want.

So this is how I am doing it right now, because everything else I tried so far
was even worse. In the Anna project we use the `vendor` directory. To manage
dependencies we use https://github.com/Masterminds/glide. When you want to add a
dependency do the following.
```
GOPATH=$(pwd)/.workspace/ glide get github.com/alicebob/miniredis
```

When you want to fetch a specific dependency, e.g. using a commit hash, do the
following.
```
GOPATH=$(pwd)/.workspace/ glide get github.com/alicebob/miniredis#10ddf01f45bee3c40d1af5dbaad9aa71e6f20835
```

When you want to install dependencies listed in the `glide.lock` file into the
`vendor/` directory, do the following.
```
GOPATH=$(pwd)/.workspace/ glide install
```

If the vendored dependencies are prepared we commit the changes and are done.
