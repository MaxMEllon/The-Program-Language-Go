#!/usr/bin/env bash

rm -rf *github.com*
go build -o out
./out https://github.com
rm -rf out

open ./github.com/index.html
