package popcount

var pc [256]byte

func init() {
	for i := range pc {

		pc[i] = pc[i/2] + byte(i&1)
	}

}


func PopCount(x uint64) int {

 var i uint = 0
 y := pc[byte(x)]
 for i < 9 {
  y += pc[byte(x>>(i*8))]
 }

 return int(y)
}
