package main

import (
	"fmt"
)

func main() {

	s := "12345678910"

	ret := comma(s)
	fmt.Print(ret) //12,345,678,910

}

func comma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	fmt.Print(s[:n-3])
	return comma(s[:n-3]) + "," + s[n-3:]
	// s[:n-3] = 12345678
	// s[n-3:] = 910
}
