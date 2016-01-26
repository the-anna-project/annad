# development
This document describes the development environment and its processes.

### setup
At first, clone the repository.
```
git clone git@github.com:xh3b4sd/anna.git
```

Fetching all dependencies works with this.
```
make goget
```

Just running the code without installing the binary can be achieved using the
go tools.
```
make gorun
```

Cleanup the workspace can be done with this.
```
make goclean
```

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
