# base image
FROM golang:1.12.4-alpine

# add maintainer info
LABEL maintainer="Simon Artner <simon.artner@gmail.com>"

# go env variables
ENV GO111MODULE on

# system dependecies
RUN apk add git

# set working directory
RUN mkdir -p $GOPATH/src/github.com/Boilertalk/kubernetes-key-rotator
WORKDIR $GOPATH/src/github.com/Boilertalk/kubernetes-key-rotator

# copy everything
COPY . $GOPATH/src/github.com/Boilertalk/kubernetes-key-rotator/

# download all dependencies
RUN go get -d -v ./...

# install the package
RUN go install -v ./...

# run server
CMD ["kubernetes-key-rotator"]
