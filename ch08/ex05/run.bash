#!/usr/bin/env bash

go build ./parallel/mandelbrot_parallel.go
go build ./synchronized/mandelbrot_synchronized.go

echo "sync"
time ./mandelbrot_synchronized > /dev/null 2>&1

echo "-----------------"

echo "parallel"
time ./mandelbrot_parallel > /dev/null 2>&1

rm -rf mandelbrot_parallel mandelbrot_synchronized
