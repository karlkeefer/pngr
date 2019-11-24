FROM postgres:12.1

RUN apt update
RUN apt install -y git-all golang-go

RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate && cp /root/go/bin/migrate /bin/migrate

# copy migration script to be run at startup
COPY run-migrations.sh /docker-entrypoint-initdb.d
