FROM postgres:12-alpine

# install pip && pgmigrate
RUN apk add --no-cache git make musl-dev go

RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

# copy migrations into the image
COPY migrations /docker-entrypoint-initdb.d/migrations

# copy migration script to be run at startup
COPY run-migrations.sh /docker-entrypoint-initdb.d