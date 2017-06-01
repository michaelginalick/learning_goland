package popcount
import (
  "fmt"
)

var pc [256]byte

func init() {
  for i := range pc {
    fmt.Println(pc[i])
    pc[i] = pc[i/2] + byte(i&1)
  }

}


func PopCount(x uint64) int {

 var i uint = 0
 y := pc[byte(x)]
 for i < 9 {
  y += pc[byte(x&(i*8))]
  // y += pc[byte(x&(x-1))]
 }

 return int(y)
}
