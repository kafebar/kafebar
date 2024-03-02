#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

go run github.com/golang-migrate/migrate/v4/cmd/migrate@latest create \
    -ext sql \
    -dir $SCRIPT_DIR/../db/migrations \
    -seq $1
