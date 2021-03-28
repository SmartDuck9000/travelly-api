FROM golang:1.15-alpine AS build

WORKDIR $GOPATH/src/github.com/SmartDuck9000/travelly-api/services/full_info_service
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT $GOPATH/bin/full_info_service
EXPOSE 5005