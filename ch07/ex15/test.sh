#!/usr/bin/env bash

SCRIPT_DIR=$(cd $(dirname $0);pwd)
cd $SCRIPT_DIR/eval && go test -v
