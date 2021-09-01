#!/bin/bash

migrate -path="/docker-entrypoint-initdb.d/schema/" -database="postgres://postgres@/postgres?host=/var/run/postgresql/" up
