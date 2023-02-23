FROM golang:1.20.1 as dev
WORKDIR /root
RUN go install github.com/cortesi/modd/cmd/modd@latest
COPY go.* ./
RUN go mod download
COPY . .
CMD modd -f worker.modd.conf

FROM golang:1.20.1 as build
WORKDIR /root
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o workerbin ./cmd/worker/worker.go

FROM gcr.io/distroless/static as prod
COPY --from=build /root/workerbin /workerbin
CMD ["/workerbin"]
