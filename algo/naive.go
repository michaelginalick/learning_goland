package main

import(
  "fmt"
)


func main() {
  a := 10
  b := 20
  z := 0


  for a > 0 {
    if a%2 == 1 {
      z = z + b
    }
    b = b << 1
    a = a >> 1
  }

  fmt.Println(z)
}


