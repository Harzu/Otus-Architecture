#! /usr/bin/env bash

export HW_HTTP_BIND_PORT=8000
export HW_DBSN='host=localhost port=5432 user=arch_homework password=arch_homework dbname=arch_homework sslmode=disable connect_timeout=5 binary_parameters=yes'
export HW_REDIS_ENABLED=true
export HW_REDIS_URL=localhost:6379
export HW_REDIS_PREFIX="products"

(cd ../migrations && goose postgres "$HW_DBSN" up)

go run ../