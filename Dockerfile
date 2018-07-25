# This dockerfile builds a production-ready version of the app

# first build a standalone binary
FROM golang:1.10-alpine
WORKDIR /root/
COPY back .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# then build a production build of front-end
FROM node:10.7-alpine
WORKDIR /root/
COPY front .
RUN npm install
RUN npm run build -o app

# then copy the built binary and static files into a minimal image
FROM alpine:latest
WORKDIR /root/
COPY --from=0 /root/app .
COPY --from=1 /root/build /root/front

CMD ./app