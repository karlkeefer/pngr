#!/bin/bash

DIR=$(dirname "$0")

if [ -f /.dockerenv ]; then
	echo "Attempting to run migrations from within the container"
	cd /docker-entrypoint-initdb.d
	echo ""
	echo "*** MIGRATIONS ***"
	# note that we are connecting via UNIX socket here, instead of normal localhost
	# this is because the docker postgres container doesn't expose postgres "ports" during the entrypoint script phase
	migrate -path="/docker-entrypoint-initdb.d/migrations/" -database="postgres://postgres@/postgres?host=/var/run/postgresql/" up
	echo "*** DONE ***"
else
	echo "Run migrations via docker-compose..."
	docker-compose -f $DIR/../docker-compose.yml exec postgres /docker-entrypoint-initdb.d/run-migrations.sh
fi
