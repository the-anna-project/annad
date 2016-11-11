# vendoring
Dependency management in go is a joke. There is actually none. There is this
idea about vendoring but there is somehow no real tooling around that provided
by the golang gang. There are loads of information about it out there. Just
google `golang vendoring`. [There are also loads of
approaches](https://github.com/golang/go/wiki/PackageManagementTools) that try
to do it. All of them are not what you probably want.

So this is how I am doing it right now, because everything else I tried to far
was even worse. In the Anna project we use the `vendor` directory. To add or
update some package to it simply `go get` the package and move it to the vendor
directory. Note that `go get` loads packages into `$GOPATH/src/`. The vendor
directory does not follow this structure. That is why the actual package needs
to be moved and the `pkg` and `src` directories need to be removed afterwards.
Once the addition or update is done, the changes need to be committed. Done.
```
GOPATH=$(pwd)/vendor/ go get -u github.com/alicebob/miniredis
mv vendor/src/* vendor
rm -rf vendor/pkg
rm -rf vendor/src
```
