package main

import (
	"fmt"
)

func main() {

	s := "12345678910"

  ret := commaRecursive(s)
  fmt.Print(ret) //12,345,678,910

}


func commaRecursive(s string) string {

	n := len(s)
	if n < 3 {
		return s
	}

	return commaRecursive(s[:n-3]) + "," + s[n-3:]
}

func commaIterative(x string) string {
	for i := len(x)-1; i > 0; i-- {
		
		if i % len(x) == 0 {
			x = x[:i] + "," + x[i:]
		}
	}
	return x
}

