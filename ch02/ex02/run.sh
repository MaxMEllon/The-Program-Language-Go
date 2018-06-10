#!/usr/bin/env bash

SCRIPT_DIR=$(cd $(dirname $0);pwd)
go build -o $SCRIPT_DIR/bin/cf $SCRIPT_DIR/src
$SCRIPT_DIR/bin/cf 100
