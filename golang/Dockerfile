FROM golang:1.15 as dev
WORKDIR /root
RUN GO111MODULE=on go get github.com/cortesi/modd/cmd/modd
COPY . .
CMD modd

FROM golang:1.15 as build
WORKDIR /root
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM gcr.io/distroless/static as prod
COPY --from=build /root/app /app
CMD ["/app"]