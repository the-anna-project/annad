# development
This document describes the development environment and its processes.

### setup
In golang the project structure is something you need to deal with. When you
ask 5 people how to do it, you probably get six answers. Here I simply describe
my personal workflow. Setup the project as you like so it fits your workflow.
So this is how I am doing it.

###### requirements
To develop and run the projects unit test suite there is only your favorite
editor and `>=go1.5` required.

There is a redis storage implementation. Using this a running redis instance is
required. For convenience this can be run in a docker container. Obviously then
docker is required as well.

###### directory structure
In golang the `GOPATH` assumes that there is something like
`src/github.com/xh3b4sd/anna/` within your workspace. See
https://golang.org/doc/code.html. I am using the `Makefile` to execute all go
commands. Note how the `GOPATH` is set in the `Makefile`.
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

---

###### clone repository
Now I am going into the working directory that holds the source code of the
Anna project.
```
mkdir -p ~/projects/private/anna/
cd ~/projects/private/anna/
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
`.workspace` directory was extended with the dependencies.
```
~/projects/private/anna
├── ...
└── .workspace
    ├── github.com
    │   ├── ...
    ├── golang.org
    │   └── ...
    └── gopkg.in
        └── ...
```

---

###### run tests
Running all tests works that way.
```
make gotest
```

---

###### remove dependencies
Cleanup the workspace can be done with this.
```
make goclean
```

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
