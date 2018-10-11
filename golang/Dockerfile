# This dockerfile builds a dev version of the backend
# It only really works as expected when run via docker-compose in the parent directory

FROM golang:1.11

WORKDIR /pngr

RUN go get github.com/cortesi/modd/cmd/modd

COPY . .

CMD modd