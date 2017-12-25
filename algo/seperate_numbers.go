package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	arr := make([]int, input)

	for i := input - 1; i < input; i++ {
		var str int
		fmt.Scan(&str)
		arr[i] = str
	}

	for i, _ := range arr {
		beautiful(arr[i])
	}
}

func beautiful(arr int) int {

	var i int
	newArr := make([]int, 100)
	for arr != 0 {
		newArr[i] = arr % 10
		arr = arr / 10
		i++
	}

	n := i
	i = 0
	for i < n {
		i++
	}

	return arr
}
