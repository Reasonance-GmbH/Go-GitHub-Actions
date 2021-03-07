FROM golang:latest

RUN mkdir project

WORKDIR /go/project
ENV GOPATH=/go/bin
ENV GO111MODULE=on
COPY go.mod .
COPY src ./src

# import dependency
RUN git clone https://github.com/manuel-lang/Go-GitHub-Actions-Dependency.git --branch v5 /usr/local/go/src/manuel-lang/Go-GitHub-Actions-Dependency

RUN go build src/main.go
CMD ["./main"]