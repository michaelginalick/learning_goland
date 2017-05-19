package main
import "fmt"

func main() {
  var numbers []int = []int{5,4,2,3,1,0}
  fmt.Println("Our list if numbers is:", numbers)
  bubble_sort(numbers)
  fmt.Println("After numbers are sorted:", numbers)
}

func bubble_sort(numbers []int) {
  var N int = len(numbers)
  var i int

  for i = 0; i< N; i++ {
    if !sweep(numbers, i) {
      return
    }
  }
}


func sweep(numbers []int, previousPasses int) bool {
  var N int = len(numbers)
  var firstIndex int = 0
  var secondIndex int = 1
  var didSwap bool = false

  for secondIndex < (N - previousPasses) {
    var firstNumber int = numbers[firstIndex]
    var secondNumber int = numbers[secondIndex]

    if firstNumber > secondNumber {
      // swap the numbers
      numbers[firstIndex] = secondNumber
      numbers[secondIndex] = firstNumber
      didSwap = true
    }

    firstIndex++
    secondIndex++
  }
  return didSwap
}
