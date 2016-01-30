FROM golang:1.4.3-alpine

RUN apk add --update git

WORKDIR /myapp
COPY . /myapp
RUN go get github.com/gorilla/handlers

EXPOSE 8080
CMD go run say_hello.go
