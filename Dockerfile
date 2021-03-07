FROM golang:latest

RUN mkdir project

WORKDIR /go/project
ENV GOPATH=/go/bin
ENV GO111MODULE=on
COPY go.mod .
COPY src ./src
RUN go build src/main.go src/operations.go
CMD ["./main"]