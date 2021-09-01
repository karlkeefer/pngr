#!/bin/bash

echo ""
echo "*** MIGRATIONS ***"
migrate -path="/docker-entrypoint-initdb.d/schema/" -database="postgres://postgres@/postgres?host=/var/run/postgresql/" up
echo "*** DONE ***"
