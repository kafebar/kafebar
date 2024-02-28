#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

DIST_DIR=$SCRIPT_DIR/../dist
API_DIR=$SCRIPT_DIR/../api
UI_DIR=$SCRIPT_DIR/../ui

rm -rf $DIST_DIR

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $DIST_DIR/main $API_DIR/src/cmd/main.go

cd $UI_DIR
pnpm install
pnpm build

cp -r $UI_DIR/dist $DIST_DIR/ui
cp $SCRIPT_DIR/../prod.env $DIST_DIR/.env

ssh jakub12134@34.0.246.58 'sudo killall main && rm -rf /home/jakub12134/dist'

scp -pr $DIST_DIR/ jakub12134@34.0.246.58:/home/jakub12134/dist

ssh jakub12134@34.0.246.58  'sudo \
    PORT=443 \
    UI_PATH=./dist/ui \
    TLS_KEY_FILE=/etc/letsencrypt/live/kafebar.pl/privkey.pem \
    TLS_CERT_FILE=/etc/letsencrypt/live/kafebar.pl/fullchain.pem \
    ./dist/main \
    </dev/null \
    >kafebar.log \
    2>&1 \
    &'