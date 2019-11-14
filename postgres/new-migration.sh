#!/bin/bash

# this janky script attempts to automatically create a new migration for you
# it takes a single argument which is used to describe the migration

DIR=$(dirname "$0")

function usage() {
	echo "You must supply a migration description as an argument"
	echo "e.g. ./new-migration.sh add_unique_constraint_to_user_email"
}

function die() {
	usage 
	exit 1
}

[ "$#" -eq 1 ] || die 

docker-compose -f $DIR/../docker-compose.yml exec postgres migrate create -dir="/docker-entrypoint-initdb.d/migrations" -ext="sql" $1