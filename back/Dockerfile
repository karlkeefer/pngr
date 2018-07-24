# first build a standalone binary
FROM golang:1.10-alpine
WORKDIR /root/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# then copy the built binary into a minimal image
FROM alpine:latest
WORKDIR /root/
COPY --from=0 /root/main .
CMD ["./main"]