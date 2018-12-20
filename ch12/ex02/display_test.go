// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package display

import (
	"fmt"
	"testing"
)

// This test ensures that the program terminates without crashing.
func TestArrayMap(t *testing.T) {
	key := [3]int{1, 2, 3}
	type MA map[[3]int]string
	ma := make(MA)
	ma[key] = "hoge"
	Display("ma", ma)
}

func TestNestedArrayMap(t *testing.T) {
	key := [1][1]int{[1]int{1}}
	type NMA map[[1][1]int]string
	nma := make(NMA)
	nma[key] = "poge"
	Display("nma", nma)
}

func TestStructMap(t *testing.T) {
	key := struct {
		key string
	}{
		key: "hoge",
	}
	type MS map[struct{ key string }]string
	nma := make(MS)
	nma[key] = "poge"
	Display("nma", nma)
}

func TestNestedStructMap(t *testing.T) {
	type Key struct {
		key struct {
			childKey string
		}
	}

	var key = Key{
		struct {
			childKey string
		}{
			childKey: "hoge",
		},
	}
	type MS map[struct{ key struct{ childKey string } }]string
	nma := make(MS)
	nma[key] = "poge"
	Display("nma", nma)
}

func Test再帰的定義(t *testing.T) {
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	err := Display("c", c)
	fmt.Println(err)
	if err == nil {
		t.Error("expected return error")
	}
}
