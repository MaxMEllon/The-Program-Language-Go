#!/usr/bin/env bash

#!/bin/bash

go build -o mandelbrot ./mandelbrot_parallel.go

for i in $(seq 1 16); do
  echo "GOMAXPROCS=$i"
  export GOMAXPROCS=$i
  time ./mandelbrot > "madelbrot.$i.png"
done
