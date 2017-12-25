package main

import (
	"fmt"
)

func main() {
	r_string := "Hello World"

	c := pallindrone(r_string)
	fmt.Println(c)
}

func pallindrone(r_string string) bool {

	rune := []rune(r_string)

	for i, j := 0, len(rune)-1; i < len(rune)/2; i, j = i+1, j-1 {
		if rune[i] != rune[j] {
			return false
		}
	}

	return true

}
