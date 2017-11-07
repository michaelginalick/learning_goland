package main
import(
    "fmt"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT

    var x1, v1, x2, v2 int
    
    _, err := fmt.Scanf("%d %d %d %d", &x1, &v1, &x2, &v2)
    if err != nil {
        fmt.Println("YES")
    }
    
    x := evaluateSteps(x1, v1, x2, v2) 
    fmt.Println(x)
}
    
func evaluateSteps(x1, v1, x2, v2 int) string {
    if v2 > v1 {
        return "NO"
    }
    
    if x1 < x2 && v1 > v2 {
        xDiff := x1-x2
        vDiff := v1-v2
        
        if xDiff % vDiff == 0 {
            return "YES"
        } else {
            return "NO"
        }
    }
    
    return "NO"
}
