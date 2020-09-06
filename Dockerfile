# inherit from the Go official Image
FROM golang:latest

# set a workdir inside docker
WORKDIR /go/src/something

# copy . (all in the current directory) to . (WORKDIR)
COPY . .

# run a command - this will run when building the image
RUN go build -o main

# run a command when running the container
CMD ./main