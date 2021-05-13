FROM ubuntu:20.04 AS build
LABEL maintainer="sonbyj01@gmail.com"

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

CMD ["GOPATH=$(pwd)", "go run main.go"]

EXPOSE 8080/TCP