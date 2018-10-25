#!/usr/bin/env bash

for i in 1 10 100 1000 10000 100000 1000000 10000000; do
  echo "pipe num: $i"
  go run main.go -p $i
done
