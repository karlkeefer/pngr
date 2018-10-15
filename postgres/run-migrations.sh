#!/bin/bash

DIR=$(dirname "$0")

if [ -f /.dockerenv ]; then
	echo "Attempting to run migrations from within the container"
	cd /docker-entrypoint-initdb.d
	pgmigrate -t latest migrate
	echo "Done!"
else
	echo "Run migrations via docker-compose..."
	docker-compose -f $DIR/../docker-compose.yml exec postgres /docker-entrypoint-initdb.d/run-migrations.sh
fi
