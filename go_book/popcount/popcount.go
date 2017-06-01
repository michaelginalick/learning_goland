package popcount

var pc [256]byte

func init() {
	for i := range pc {

		pc[i] = pc[i/2] + byte(i&1)
	}

}


func PopCount(x uint64) int {

 for i uint = 0; i < 9; i++ {
   x += int(pc[byte(x>>(i*8))])
 }
 return int(pc[byte(x)])
}
