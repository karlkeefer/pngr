#!/bin/bash

# this janky script attempts to automatically create a new migration for you
# it takes a single argument which is used to describe the migration

function usage() {
	echo "You must supply a migration description as an argument"
	echo "e.g. ./new-migration.sh add_unique_constraint_to_user_email"
}

function die() {
	usage 
	exit 1
}

[ "$#" -eq 1 ] || die 

# get last migration filename
DIR=$(dirname "$0")
LAST_MIGRATION=$(ls "$DIR/migrations" | tail -n 1)

# pull out the version number
OLD_VERSION="${LAST_MIGRATION:1:4}"

# increment it and leftpad
NUM=$(printf "%04d" $(($OLD_VERSION + 1)))

# build new filename
NEW_MIGRATION=$(printf "V%s__%s.sql" $NUM $1)

# create the file with the `ops` insert already populated
echo "-- Your migration goes here --

INSERT INTO ops (op) VALUES('migration $NEW_MIGRATION');" > $DIR/migrations/$NEW_MIGRATION