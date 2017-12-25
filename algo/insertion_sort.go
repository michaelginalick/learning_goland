package main

import (
	"fmt"
)

func main() {

	arr := []int{31, 41, 59, 26, 61, 58}
	insertionSort(arr)
	fmt.Println(arr)
}

//[31 26 41 58 59 61]
func insertionSort(arr []int) []int {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1

		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i = i - 1
		}
		arr[i+1] = key
	}

	return arr
}
