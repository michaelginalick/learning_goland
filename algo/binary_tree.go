package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {

	s := []int{54, 99, 23, 32, 2, 90, 31, 1, 9999, 10}
	root := new(tree)
	root = nil
	newTree := insert(s, root)
	newTree = lookup(newTree, s[4])
	fmt.Println(newTree)
}

func insert(values []int, root *tree) *tree {
	for _, v := range values {
		root = add(root, v)
	}
	return root
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


func lookup(t *tree, value int) *tree {
	if t == nil {
		return t
	}
	if t.value == value {
		return t
	} else if value < t.value {
		return lookup(t.left, value)
	} else {
		return lookup(t.right, value)
	}
}