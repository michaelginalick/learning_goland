package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := []int{5, 4, 2, 3, 1, 0, 90, 123, 45, 345, 32, 21, 22, 33, 44, 56, 76, 788, 110, 99, 23}
	fmt.Println(numbers)
	quick_sort(numbers)
	fmt.Println(numbers)

}

func quick_sort(numbers []int) []int {

	if len(numbers) < 2 {
		return numbers
	}

	left, right := 0, len(numbers)-1
	//pick a pivot
	pivotIndex := rand.Int() % len(numbers)

	//move the pivot to the right
	swap(numbers, pivotIndex, right)

	//pile element smaller than the pivot to the left
	for i := range numbers {
		if numbers[i] < numbers[right] {
			swap(numbers, i, left)
			left++
		}
	}

	//place the pivot after the smaller element
	swap(numbers, left, right)

	//go down the rabbit hole
	quick_sort(numbers[:left])
	quick_sort(numbers[left+1:])

	return numbers
}

func swap(numbers []int, i, j int) []int {
	var temp int

	temp = numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = temp

	return numbers
}
