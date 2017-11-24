package main
import "fmt"

func main() {
    var str string
    
    fmt.Scan(&str)
    rune := []rune(str)
    
    for i := 0; i < len(rune)-1; i++ {
        if string(rune[i]) == string(rune[i+1]) {
            rune = append(rune[:i], rune[i+1:]...)
            rune = append(rune[:i], rune[i+1:]...)
            i = -1
        }
    }
    newStr := string(rune)
    if len(newStr) == 0 {
        fmt.Print("Empty String")
    } else {
           fmt.Print(newStr)   
    }
}
