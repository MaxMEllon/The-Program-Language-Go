#!/usr/bin/env bash

SCRIPT_DIR=$(cd $(dirname $0);pwd)
go test -bench=.
