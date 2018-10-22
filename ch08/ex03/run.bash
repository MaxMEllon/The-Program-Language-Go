#!/usr/bin/env bash

trap "kill 0" EXIT

go run ./reverb.go &

go run ./netcat.go
