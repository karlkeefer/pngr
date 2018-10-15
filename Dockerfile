# This dockerfile builds a production-ready version of the app

# first build a standalone binary
# TODO: consider basing these on the development dockerfile
# 		so that versions automatically match
FROM golang:1.11
WORKDIR /pngr/golang
COPY golang .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# then build a production build of front-end
# TODO: consider basing these on the development dockerfile
# 		so that versions automatically match
FROM node:10.7
WORKDIR /pngr/react
COPY react .
RUN npm install
RUN npm run build -o app

# then copy the built binary and static files into a minimal image
FROM alpine:latest
WORKDIR /root
COPY --from=0 /pngr/golang/app .
COPY --from=1 /pngr/react ./front

CMD ./app