#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

source $SCRIPT_DIR/../.env

go run -tags postgres github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
    -path $SCRIPT_DIR/../db/migrations \
    -database $POSTGRES_CONNSTRING \
    up
