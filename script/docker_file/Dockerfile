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
    && rm -rf docker docker-latest.tgz

# install expect
RUN apt-get update
RUN apt-get -y install tcl tk expect
