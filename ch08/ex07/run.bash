#!/usr/bin/env bash

rm -rf *golang.org*
go build -o out
./out https://golang.org
rm -rf out

open ./golang.org/index.html
