#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

cd  $SCRIPT_DIR/../api
go run github.com/joho/godotenv/cmd/godotenv@latest -f $SCRIPT_DIR/../.env go run kafebar/cmd/main.go