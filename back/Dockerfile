# This dockerfile builds a dev version of the backend
# It only really works as expected when run via docker-compose in the parent directory

FROM golang:1.10
WORKDIR /root/
RUN go get github.com/codegangsta/gin
CMD gin --port 5000 --appPort 3030 run /root/main.go