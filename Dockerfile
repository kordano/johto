FROM golang:alpine
MAINTAINER Konrad Kühne <konrad.kuehne@lambdaforge.io>

# Adapted from https://github.com/callicoder/go-docker-compose/blob/master/Dockerfile
RUN apk add git

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/kordano/johto

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

CMD ["johto"]
