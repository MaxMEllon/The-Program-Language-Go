package main

import (
	"fmt"
	"sort"
)

func uniq(stringList []string) []string {
	if len(stringList) == 0 {
		return make([]string, len(stringList))
	}

	sort.Slice(stringList, func(i, j int) bool {
		return stringList[i] < stringList[j]
	})

	result := stringList[:1] // first item is a always uniq.

	for _, s := range stringList[1:] {
		if result[len(result)-1] != s {
			result = append(result, s)
		}
	}
	return result
}

func main() {
	fmt.Println(uniq([]string{
		"vim",
		"emacs",
		"vim",
		"emacs",
		"vim",
		"emacs",
		"vim",
		"emacs",
		"vim",
		"emacs",
		"vim",
	}))
}
