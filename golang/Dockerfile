# This dockerfile builds a dev version of the backend
# It only really works as expected when run via docker-compose in the parent directory

FROM golang:1.10

WORKDIR /go/src/github.com/karlkeefer/pngr/golang

RUN go get github.com/golang/dep/...
RUN go get github.com/cortesi/modd/cmd/modd

COPY . .
RUN dep ensure -vendor-only

CMD modd