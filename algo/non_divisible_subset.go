package main
import "fmt"


func main() {
     var numCount, divisor int

    fmt.Scanf("%d %d", &numCount, &divisor)
    remainder := make([]int, numCount)

    for i := 0; i < numCount; i++ {
      var a int
      fmt.Scanf("%d",&a)
      remainder[a % divisor]++
    }
    
    
    var max int
    
    for i := 0; i <= divisor/2; i++ {
        if i == 0 || i == divisor - i {
            if remainder[i] >= 1 {
                max += 1
            }
        } else {
            
            if remainder[i] > remainder[divisor-i] {
                max += remainder[i]
            } else {
                max += remainder[divisor-i]
            }
        }
    }
        
    fmt.Println(max)
}
