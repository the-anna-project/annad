# contributing
Here is described how to contribute to the Anna project.

### pull requests
All changes affecting the Anna project MUST be provided in form of a proper PR.
What exactly "proper" means evolves over time and is not written in stone. Some
important points are listed below. To help setting up a PR the pull request
template [PULL_REQUEST_TEMPLATE.md](/.github.com/PULL_REQUEST_TEMPLATE.md) will
automatically be provided within the pull request's description.

### labels and assignees
Please make sure sufficient labels are applied and you have assigned yourself
to your pull request. If there is no sufficient label for your PR, please file
an issue.

### commits
Pull requests are only accepted and merged when there is only one commit to be
merged. This means contributers need to squash their commits before. This can
be done with the following command.
```
git rebase -i master
```

### docs
Pull requests are only accepted and merged when there is proper documentation.

Conceptual documentation needs to be provided here: https://github.com/xh3b4sd/anna/tree/master/doc/concept

Process documentation needs to be provided here: https://github.com/xh3b4sd/anna/tree/master/doc/diary

Code documentation needs to be provided within the code: http://blog.golang.org/godoc-documenting-go-code

### tests
Pull requests are only accepted and merged when there are proper tests. Make
sure there are tests where it makes sense and all tests pass reliably. No pull
request is going to be merged as long as there are tests failing or flapping.
It is not mandatory to increase code coverage as long as there are sufficient
tests.

#### branches
When colaborating and creating branches within the core repository
https://github.com/xh3b4sd/anna branches should be deleted ones they are
merged. This prevents loosing the overview and keeps the repository in a clean
state branch wise.
