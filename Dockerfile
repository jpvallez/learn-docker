# inherit from the Go official Image
FROM golang:1.8

# set a workdir inside docker
WORKDIR /go/src/server

# copy . (all in the current directory) to . (WORKDIR)
COPY . .

# run a command - this will run when building the image
RUN go build -o main

# the port we wish to expose
EXPOSE 8080

# run a command when running the container
CMD ./main