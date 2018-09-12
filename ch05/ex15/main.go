package main

import (
	"fmt"
	"math"
)

const (
	MIN = math.MinInt64
	MAX = math.MaxInt64
)

func expectedOverOneArgs(vals []int64) error {
	if len(vals) == 0 {
		return fmt.Errorf("expected over one args")
	}
	return nil
}

func max(vals ...int64) (result int64, err error) {
	result = MIN
	if err := expectedOverOneArgs(vals); err != nil {
		return 0, err
	}
	for _, v := range vals {
		if result < v {
			result = v
		}
	}
	return
}

func min(vals ...int64) (result int64, err error) {
	if err := expectedOverOneArgs(vals); err != nil {
		return 0, err
	}
	result = MAX
	for _, v := range vals {
		if result > v {
			result = v
		}
	}
	return
}

func main() {
	fmt.Println(min(0, 1, 2, 4, 5, 6))
	fmt.Println(max(0, 1, 2, 4, 5, 6))
}
