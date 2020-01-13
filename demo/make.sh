#!/bin/sh

export DATABASE=postgres://postgres:postgres@localhost/postgres?sslmode=disable
psql -h localhost --username="postgres" -d postgres -a -f demo/data.sql
