package main

import "fmt"

func main() {
	var pattern_string string
	var count, sum, sum2 int

	fmt.Scanf("%s", &pattern_string)
	fmt.Scanf("%d", &count)

	for i := 0; i < len(pattern_string); i++ {
		if string(pattern_string[i]) == "a" {
			sum++
		}
	}

	sum *= (count / len(pattern_string))

	for i := 0; i < (count % len(pattern_string)); i++ {
		if string(pattern_string[i]) == "a" {
			sum2++
		}
	}

	fmt.Print(sum + sum2)

}
