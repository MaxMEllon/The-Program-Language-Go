package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

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
	res, err := topoSort(prereqs)
	if err != nil {
		panic("invalid")
	}
	for i, course := range res {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func containsInSlice(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string, parents []string) error

	visitAll = func(items []string, parents []string) error {
		for _, item := range items {
			// fmt.Printf("%v\n", parents)
			if containsInSlice(parents, item) {
				return fmt.Errorf("loop")
			}

			if !seen[item] {
				seen[item] = true
				parents = append(parents, item)
				err := visitAll(m[item], parents)
				if err != nil {
					return err
				}
				order = append(order, item)
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	err := visitAll(keys, []string{})
	return order, err
}
