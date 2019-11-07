FROM postgres:12

# install pip && pgmigrate
RUN apt-get update && apt-get install -y golang git \
	&& go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

# copy migrations into the image
COPY migrations /docker-entrypoint-initdb.d/migrations

# copy migration script to be run at startup
COPY run-migrations.sh /docker-entrypoint-initdb.d