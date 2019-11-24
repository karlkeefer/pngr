FROM postgres:12.1

RUN apt-get update && apt-get install -y golang git

RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate && cp /root/go/bin/migrate /bin/migrate

# copy migration script to be run at startup
COPY run-migrations.sh /docker-entrypoint-initdb.d