# debugging
Debugging is an important topic and sometimes extremely critical when it comes
to tackle nasty bugs. Here we document some debugging techniques that should
help in certain situations.

### pprof
`pprof` is a useful tool build into the go tools. It can be used to gather
interesting information about e.g. memory and CPU usage. Lets say there is
something fishy within your code. To debug this we can make use of `pprof`
using the following command. Note that the command is very specific to the Anna
project, but you should get the idea. We create a test binary and a memory and
CPU profile.
```
GOPATH=$(pwd)/.workspace/ go generate ./... && GOPATH=$(pwd)/.workspace/ go test ./index/clg/collection -v -run Test_FeatureSet_GetFeaturesByCountFeatureSet_Expected -cpuprofile cpu.out -memprofile mem.out
```

Having the binary and profiles in place we can visualize the programs memory
and CPU usage like that. The created PDFs in this case can simply be viewed in
your browser.
```
go tool pprof -pdf collection.test cpu.out > id_factory_cpu.pdf
go tool pprof -pdf collection.test mem.out > id_factory_mem.pdf
```

Note that e.g. on ubuntu 14.04 you can install the required graphing
dependencies using the following command.
```
sudo apt-get install graphviz
```

For more information consider the following links.
- https://blog.golang.org/profiling-go-programs
- https://golang.org/cmd/go/#hdr-Description_of_testing_flags
