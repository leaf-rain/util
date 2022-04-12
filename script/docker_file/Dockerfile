# Base image: https://hub.docker.com/_/golang/
FROM golang:1.17
USER root
# Install golint
ENV GOPATH=/go PATH=${GOPATH}/bin:$PATH GO111MODULE=on GOPROXY="https://proxy.golang.org,direct"
RUN go get -u golang.org/x/lint/golint

# install docker
RUN curl -O https://get.docker.com/builds/Linux/x86_64/docker-latest.tgz \
    && tar zxvf docker-latest.tgz \
    && cp docker/docker /usr/local/bin/ \
    && rm -rf docker docker-latest.tgz && mkdir -p /etc/docker
RUN echo $'{"registry-mirrors": ["https://q3u0qv3z.mirror.aliyuncs.com"],"insecure-registries":["192.168.2.200:5000"]}' > /etc/docker/daemon.json

# install expect
RUN apt-get update && apt-get -y install tcl tk expect
