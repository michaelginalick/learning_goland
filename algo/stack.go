package main
import "fmt"

type Stack struct {
  slice []int
}

func main() {

  var s *Stack = new(Stack)
	s.Push(1)
	s.Push(99)
	s.Push(102)
	s.Push(1000)
	fmt.Println(s)
  fmt.Println(s.Peek())
  fmt.Println(s.Pop())
  fmt.Println(s)

}

func (s *Stack) Push(i int) {
	s.slice = append(s.slice, i)
}


func (s *Stack) Peek() int {
	if len(s.slice) == 0 {
		return 0
	}
  return s.slice[len(s.slice)-1]
}

func (s *Stack) Size() int {
	return len(s.slice)
}

func (s *Stack) String() string {
	return fmt.Sprint(s.slice)
}

func (s *Stack) Pop() int {
  ret := s.Peek()
	s.slice = s.slice[0:len(s.slice)-1]
	return ret
}
