#!/usr/bin/env bash

go build -o chat ./chat.go

go build -o netcat ./netcat.go

echo "他のターミナルで netcat を起動"

./chat
