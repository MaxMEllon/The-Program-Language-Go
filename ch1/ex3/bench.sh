#!/usr/bin/env bash

for end in 5000 10000 15000 20000 25000; do
  args=`ruby -e "puts 'a ' * $end"`
  echo $end
  echo "bad"
  bash -c "time -p go run bad.go $args  2>/dev/null"
  echo "good"
  bash -c "time -p go run good.go $args 2>/dev/null"
done


