FROM golang:1.15-alpine AS build

WORKDIR $GOPATH/src/github.com/SmartDuck9000/travelly-api/backend_1
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT $GOPATH/bin/backend_1
EXPOSE 5001