#!/bin/sh

export DATABASE=postgres://postgres:postgres@localhost/postgres?sslmode=disable
psql -h localhost --username="postgres" -d postgres -a -f data/clear.sql
migrate -database ${DATABASE} -path migrations down
migrate -database ${DATABASE} -path migrations up

