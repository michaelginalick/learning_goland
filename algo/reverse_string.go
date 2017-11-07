package main

import (
	"fmt"
)

func main() {
	r_string := "Hello World"
	rune := []rune(r_string)

	for i, j := 0, len(rune)-1; i < len(rune)/2; i, j = i+1, j-1 {
		rune[i], rune[j] = rune[j], rune[i]
	}

	fmt.Println(string(rune))
}
