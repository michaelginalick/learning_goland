package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)

	arr := make([]int, d)

	for i := 0; i < len(arr); i++ {
		var a int
		fmt.Scan(&a)
		arr[i] = a
	}

	var numOfJumps, i int

	for {
		if i+2 < len(arr) && arr[i+2] == 0 {
			i += 2
		} else if i+1 < len(arr) {
			i++
		} else {
			break
		}

		numOfJumps++
	}

	fmt.Print(numOfJumps)
}
