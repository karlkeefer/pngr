# This dockerfile builds a dev version of the backend
# It only really works as expected when run via docker-compose in the parent directory

FROM golang:1.10

WORKDIR /go/src/github.com/karlkeefer/pngr/golang

RUN go get github.com/codegangsta/gin
RUN go get github.com/golang/dep/...

COPY . .
RUN dep ensure -vendor-only

# make sure it builds even though we are going to rely on gin
RUN go build main.go

CMD gin --godep --port=5000 --appPort=5050 run /go/src/github.com/karlkeefer/pngr/golang/main.go