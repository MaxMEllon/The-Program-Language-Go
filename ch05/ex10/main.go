package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":      {"discrete math"},
	"databases":            {"data structures"},
	"discrete math":        {"intro to programming"},
	"formal languages":     {"discrete math"},
	"networks":             {"operating system"},
	"operating system":     {"data structures", "computer organization"},
	"programming language": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, val := range items {
			if !seen[val] {
				seen[val] = true
				visitAll(m[val])
				order = append(order, val)
			}
		}
	}

	for item := range m {
		visitAll([]string{item})
	}
	return order
}
