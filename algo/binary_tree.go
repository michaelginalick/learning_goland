package main

import(
  "fmt"
)


type tree struct {
	value int
	left, right *tree
}



func main() {

  s := []int{54, 23, 12, 2, 0, 3, 1, 9, 10, 99}
  x := insert(s)
	fmt.Println(x.left.left.left)
}



func insert(values []int) *tree {
	var root *tree
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
