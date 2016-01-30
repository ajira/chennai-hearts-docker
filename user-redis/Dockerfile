FROM golang:1.4.3-alpine

RUN apk add --update git

WORKDIR /myapp

COPY build.sh /myapp/build.sh
RUN ./build.sh

COPY . /myapp
RUN go build .
CMD go run user_register.go
