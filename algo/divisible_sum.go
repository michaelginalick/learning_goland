package main

import "fmt"

func main() {
	var length, divisor int

	fmt.Scan(&length, &divisor)

	arr := make([]int, length)

	for i := 0; i < len(arr); i++ {
		var a int
		fmt.Scan(&a)
		arr[i] = a
	}

	var max int

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if (arr[i]+arr[j])%divisor == 0 {
				max++
			}
		}
	}

	fmt.Println(max)
}
