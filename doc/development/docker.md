# docker
[Docker](https://github.com/docker/docker) provides container technology like
[Rkt](https://github.com/coreos/rkt) and [LXC](https://linuxcontainers.org). In
the Anna project we make use of docker to ship a docker image containing the
[anna](/doc/development/makefile.md#anna) and
[annactl](/doc/development/makefile.md#annactl) binaries. The docker image is
publicly available and can be pulled using the following command.

```
docker pull xh3b4sd/anna
```

See the [Dockerfile](/Dockerfile) and the [docker
repository](https://hub.docker.com/r/xh3b4sd/anna) for more information.
