FROM ubuntu:20.04 AS build
LABEL maintainer="sonbyj01@gmail.com"

# https://dev.to/setevoy/docker-configure-tzdata-and-timezone-during-build-20bk
ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt update && \
    apt install -y git golang-go && \
    git clone https://github.com/sonbyj01/ece465-final.git /final && \
    cd final/src && \
    rm -rf github.com/ golang.org/ google.golang.org/ gopkg.in/ && \
    cd .. && \
    GOPATH=$(pwd) go get github.com/markus-wa/quickhull-go && \
    GOPATH=$(pwd) go get github.com/gin-gonic/gin && \
    GOPATH=$(pwd) go get github.com/google/uuid && \
    mkdir tmp && \
    rm -rf /var/lib/apt/lists/* && \
    apt clean

#RUN "GOPATH=/final go run main.go"

EXPOSE 8080/TCP
