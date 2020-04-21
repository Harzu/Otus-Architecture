#! /usr/bin/env bash

export HW_3_HTTP_BIND_PORT=8000
export HW_3_DBSN='host=localhost port=5432 user=postgres password=password dbname=postgres_homework_3 sslmode=disable connect_timeout=5 binary_parameters=yes'

(cd ../migrations && goose postgres "$HW_3_DBSN" up)

go run ../