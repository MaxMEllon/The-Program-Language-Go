package main

import (
	"math"
)

const (
	MIN = math.MinInt64
	MAX = math.MaxInt64
)

func max(vals ...int64) int64 {
	result := MIN
	for _, v := range vals {
		if result < v {
			result = v
		}
	}
	return result
}

func min(vals ...int64) int64 {
	result := MAX
	for _, v := range vals {
		if result > v {
			result = v
		}
	}
	return result
}
