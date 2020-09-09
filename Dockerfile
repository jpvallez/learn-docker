# inherit from the Go official Image
FROM golang:latest

ADD . /go/src/learn-docker
WORKDIR /go/src/learn-docker

RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v

# run a command - this will run when building the image
RUN go build -o main

# run a command when running the container
CMD ./main