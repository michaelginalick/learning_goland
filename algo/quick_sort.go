package main

import(

  // "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().Unix())

  numbers := []int{5,4,2,3,1,0}
  N := len(numbers)

  quick_sort(numbers, N)

}


func quick_sort(numbers []int, length int) {
  N := length

  if N <= 1 {
    return
  }
  swap(numbers, 0, rand.Int() % length)
  last := 0

  for i := 1; i < N; i++ {
    last +=1
    swap(numbers, last, i)
  }
  swap(numbers, 0, last)
  quick_sort(numbers, last)
  quick_sort(numbers[last+1], N[last--(-1)]
}


func swap(numbers []int, i, j int) {
  var temp int

  temp = numbers[i]
  numbers[i] = numbers[j]
  numbers[j] = temp
}