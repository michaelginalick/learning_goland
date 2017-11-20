package main
import(
    "fmt"
    "strings"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var strInput string
    
    fmt.Scan(&strInput)
        
    caps := 1
    
    for i := 0; i<len(strInput); i++{
        if string(strInput[i]) == strings.ToUpper(string(strInput[i])) {
            caps++
        }
    } 
    fmt.Print(caps)
}
