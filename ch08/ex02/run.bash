#!/usr/bin/env bash

echo "go build -o f ./ftpd.go"
go build -o f ./ftpd.go
echo "./f"
sudo ./f
rm -rf f
