# development
This document describes the development environment and its processes.

### setup
In golang the project structure is something you need to deal with. When you
ask 5 people how to do it, you probably get six answers. Here I simply describe
my personal workflow. Setup the project as you like so it fits your workflow.
So this is how I am doing it.

###### directory structure
In golang the `GOPATH` assumes that there is something like
`src/github.com/xh3b4sd/anna/` within your workspace. See
https://golang.org/doc/code.html. I am using the `Makefile` to execute all go
commands. Note how the `GOPATH` is set in the `Makefile`.
```
$(mkdir -p .workspace/)
GOPATH := ${PWD}/.workspace/:${PWD}/../../../..:${GOPATH}
export GOPATH
```

This takes care that the workspace is properly set up for the dependencies that
`go get` fetches and that `GOPATH` itself is properly set. So the result after
the setup will be similar to this.
```
~/projects/private/anna                     # This is the second path within GOPATH and used for the anna project
├── ...
└── src
    └── github.com
        └── xh3b4sd
            └── anna
                ├── ...
                └── .workspace              # This is the first path within GOPATH and used for dependencies
                    └── ...
```

This approach ensures that there is a dedicated workspace for the anna project.
That is, `~/projects/private/anna/src/github.com/xh3b4sd/anna/.workspace/`.
This approach also ensures that the directiory structure for the anna project
itself is fulfilled, so the go tools can operate properly on it. Just following
the standard go project layout leads to the fact that all project dependencies
are lying around outside the project itself. So when I would like to clean all
dependencies, I would need to go though the directory structure and delete all
directories but that one of my own project. This is not really straight
forward, even if the bash magicians have other opinions. Other approaches like
symlinking the project into the `.workspace/` directory lead to the fact that
the go tools cannot properly operate on the project. Symlinks are prohibited by
`GOPATH`. Thus `go test ./... -cover` will not work in all cases.

So at first, I prepare the directory structure.
```
mkdir -p ~/projects/private/anna/src/github.com/xh3b4sd/anna/
```

---

###### clone repository
Now I am going into the working directory that holds the source code of the
anna project.
```
cd ~/projects/private/anna/src/github.com/xh3b4sd/anna/
```

Then, I clone the repository. Note the `.` at the end of the command.
```
git clone git@github.com:xh3b4sd/anna.git .
```

---

###### fetch dependencies
Fetching all dependencies now works with this.
```
make goget
```

Now what happened? Because of our directory structure and `Makefile` magic the
project structure looks like this.
```
~/projects/private/anna                     # This is the second path within GOPATH and used for the anna project
├── ...
└── src
    └── github.com
        └── xh3b4sd
            └── anna
                ├── ...
                └── .workspace              # This is the first path within GOPATH and used for dependencies
                    ├── ...
                    └── src
                        ├── github.com
                        │   ├── ...
                        ├── golang.org
                        │   └── ...
                        └── gopkg.in
                            └── ...
```

---

###### remove dependencies
Cleanup the workspace can be done with this.
```
make goclean
```

---

###### run tests
Running all tests works that way.
```
make gotest
```

### pull requests
All changes affecting the project MUST be provided in form of a proper PR. What
exactly "proper" means evolves over time and is not written in stone. Some
important points are listed as follows.

###### commits
Pull requests are only accepted and merged when there is only one commit to be
merged. This means contributers need to squash their commits before. This can
be done with the following command.
```
git rebase -i master
```

---

###### docs
Pull requests are only accepted and merged when there is proper documentation.

Conceptual documentation needs to be provided here: https://github.com/xh3b4sd/anna/tree/master/doc/concept

Process documentation needs to be provided here: https://github.com/xh3b4sd/anna/tree/master/doc/diary

Code documentation needs to be provided within the code: http://blog.golang.org/godoc-documenting-go-code

---

###### tests
Pull requests are only accepted and merged when there are proper tests. Make
sure there are tests where it makes sense and all tests pass reliably. No pull
request is going to be merged as long as there are tests failing or flapping.

### guidelines

###### tracability
Having insides into complex systems is key. Events going through neural
networks need to be highly comprehensible in detail. Data needs to be
collected. All information we can gather need to be visualized somehow.

---

###### plugability
A plugable architecture having decent interfaces ...

---

###### testability
Functionality needs to be guaranteed by testing actual software
implementations. Automated. Painless. Fast.

---

###### data formats
For simplicity JSON should be good enough for now.

---

###### API's
There are two forms of API's we want to care about here. Library interfaces and
network API's.

Library interfaces should always be well defined. Software packages of this
project need to have a real purpose on their own, so they can stand alone, or
be used by something else.

Network API's should always be well defined. They should simply represent
business logic implementation wrapped by some network protocol middleware.
