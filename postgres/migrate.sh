#!/bin/bash

# Usage here is the same as the normal go-migrate tool
# e.g. 
# postgres/migrate.sh up
# postgres/migrate.sh force 20191213014645
# postgres/migrate.sh up

DIR=$(dirname "$0")
docker-compose -f $DIR/../docker-compose.yml exec postgres migrate -path="/docker-entrypoint-initdb.d/migrations/" -database="postgres://postgres@/postgres?host=/var/run/postgresql/" $@
