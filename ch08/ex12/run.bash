#!/usr/bin/env bash

go build -o chat

go build -o netcat gopl.io/ch8/netcat3

echo "他のターミナルで netcat を起動"

./chat
