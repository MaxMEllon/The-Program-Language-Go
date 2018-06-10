#!/usr/bin/env bash

SCRIPT_DIR=$(cd $(dirname $0);pwd)
cd $SCRIPT_DIR/src && dep ensure && cd .. && go build -o ./bin/cf ./src
$SCRIPT_DIR/bin/cf -w -t -d 300
