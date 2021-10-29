FROM golang:1.16 as dev
WORKDIR /root
RUN GO111MODULE=on go get github.com/cortesi/modd/cmd/modd
COPY go.* ./
RUN go mod download
COPY . .
CMD modd -f worker.modd.conf

FROM golang:1.16 as build
WORKDIR /root
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o workerbin ./cmd/worker/worker.go

FROM gcr.io/distroless/static as prod
COPY --from=build /root/workerbin /workerbin
CMD ["/workerbin"]
