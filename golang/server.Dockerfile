FROM golang:1.16 as dev
WORKDIR /root
RUN go get github.com/cortesi/modd/cmd/modd
RUN go get github.com/kyleconroy/sqlc/cmd/sqlc
COPY . .
CMD modd -f server.modd.conf

FROM golang:1.16 as build
WORKDIR /root
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o serverbin ./cmd/server/server.go

FROM gcr.io/distroless/static as prod
COPY --from=build /root/serverbin /serverbin
CMD ["/serverbin"]
