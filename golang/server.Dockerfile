FROM golang:1.18 as dev
WORKDIR /root
RUN go install github.com/cortesi/modd/cmd/modd@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN go install github.com/golang/mock/mockgen@v1.6.0
COPY go.* ./
RUN go mod download
COPY . .
CMD modd -f server.modd.conf

FROM golang:1.18 as build
WORKDIR /root
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o serverbin ./cmd/server/server.go

FROM gcr.io/distroless/static as prod
COPY --from=build /root/serverbin /serverbin
CMD ["/serverbin"]
