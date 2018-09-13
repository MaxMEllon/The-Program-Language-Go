package main

import (
	"bytes"
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func trace(t *tree, depth int, buf *bytes.Buffer) {
	if t == nil {
		return
	}
	buf.WriteString(fmt.Sprintf("%*s`-%s %d\n", depth*6, " ", strings.Repeat("-", 2), t.value))
	trace(t.left, depth+1, buf)
	trace(t.right, depth+1, buf)
}

func (t *tree) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("\\\n"))
	trace(t, 0, &buf)
	return buf.String()
}

func main() {
	var t *tree
	t = add(t, 1)
	t = add(t, -1)
	t = add(t, 2)
	t = add(t, 5)
	t = add(t, -10)
	t = add(t, 10)
	t = add(t, -2)
	t = add(t, 8)
	t = add(t, -20)
	fmt.Println(t.String())
}
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
