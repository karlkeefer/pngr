FROM postgres:10.4

# install pip && pgmigrate
RUN apt-get update && apt-get install -y python-pip \
	&& pip install yandex-pgmigrate \
	&& pip install psycopg2-binary

# copy migrations into the image
COPY migrations /docker-entrypoint-initdb.d/migrations

# copy migration script to be run at startup
COPY run-migrations.sh /docker-entrypoint-initdb.d