package main
import "fmt"
import "sort"

func main() {
    var d int
    fmt.Scan(&d)
    
    m := make(map[int]int)
    
    for i := 0; i <= d; i++ {
        var a int
        fmt.Scan(&a)
        m[a]++
    }
   //place map values into array  
    values := make([]int, len(m))
    i := 0
     for k := range m {
        values[i] = m[k]
         i++
    }
    //sort values
    sort.Ints(values)
   // get max value
    x := maxValue(values)
   // find key by value 
    key, ok := mapkey(m, x)
    
    if !ok {
        panic("value does not exist in map")
    }
    
    fmt.Print(key)
}

func maxValue(a []int) int {
  x := a[0]
  for _, e := range a {
    if e > x {
      x = e
    }
  }
  return x
}


func mapkey(m map[int]int, value int) (key int, ok bool) {
  for k, v := range m {
    if v == value { 
      key = k
      ok = true
      return
    }
  }
  return
}
