# docker
[Docker](https://github.com/docker/docker) provides container technology like
[Rkt](https://github.com/coreos/rkt) and [LXC](https://linuxcontainers.org). In
the Anna project we make use of docker to ship a docker image containing the
[annad](/doc/development/makefile.md#annad) and
[annactl](/doc/development/makefile.md#annactl) binaries. See also the
[Dockerfile](/Dockerfile) and the [docker
repository](https://hub.docker.com/r/xh3b4sd/anna) for more information.

### docker pull
The docker image is publicly available and can be pulled using the following
command. Note that there is no latest tag. Check the tag recently pushed on the
docker hub.
```
docker pull xh3b4sd/anna:<tag>
```

### docker run
Once pulled, the docker image can be used to run a docker container using the
following command. Note that we execute the `annad` (server) binary here.
```
docker run xh3b4sd/anna:<tag> annad -h
```

Once pulled, the docker image can be used to run a docker container using the
following command. Note that we execute the `annactl` (client) binary here.
```
docker run xh3b4sd/anna:<tag> annactl -h
```
