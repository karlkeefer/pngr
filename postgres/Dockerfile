FROM postgres:12.1

RUN apt-get update && apt-get install -y wget

# snag a binary of golang-migrate
RUN wget -nv https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz \ 
  && tar -xzf migrate.linux-amd64.tar.gz \
  && rm migrate.linux-amd64.tar.gz \
  && cp migrate.linux-amd64 /bin/migrate

# copy migration script to be run at startup
COPY run-migrations.sh /docker-entrypoint-initdb.d