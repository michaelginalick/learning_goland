package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	s := []int{54, 23, 12, 2, 0, 3, 1, 9, 10, 99}

	Sort(s)
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		fmt.Println(root)
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
		// equivalent to return &tree{value: value}
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
