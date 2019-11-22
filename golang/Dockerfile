# This dockerfile builds a dev version of the backend
# It only really works as expected when run via docker-compose in the parent directory

FROM golang:1.13-alpine
WORKDIR /pngr/golang
RUN apk add --no-cache git make musl-dev gcc
RUN GO111MODULE=on go get github.com/cortesi/modd/cmd/modd
COPY . .
CMD modd