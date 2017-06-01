package main
import (
  "fmt"
)

var pc [256]byte

func main() {
  for i := range pc {
    fmt.Println(pc[byte(x>>(i*8))])
    pc[i] = pc[i/2] + byte(i&1)
  }

}

