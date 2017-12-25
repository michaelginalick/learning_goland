package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	size, err := strconv.Atoi(in.Text())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad array size\n")
		os.Exit(1)
	}
	a := make([]int, size)
	in.Scan()
	svals := strings.Split(in.Text(), " ")
	for _, s := range svals {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Bad array number: %s\n", in.Text())
		}
		a = append(a, i)
	}
	fmt.Println(sumArray(a))
}
func sumArray(a []int) (sum int) {
	for _, i := range a {
		sum += i
	}
	return
}
