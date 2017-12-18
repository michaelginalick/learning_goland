package main


import(
  "fmt"
)




type Student struct{
  age int
  weight int
  name string
  Next *Student
}


type Teacher struct {
  age int
  weight int
  name string
  Next *Student
}

func main() {

  kyle := Student{7, 70, "kyle", nil}
  john := Student{10, 100, "john", &kyle}

  deborah := Teacher{24, 185, "deborah", &john}

   fmt.Println("The teach is", deborah.name)
   fmt.Println("the student behind deborah is:", deborah.Next.name)
   fmt.Println("the student behind john is:", deborah.Next.Next.name)
   fmt.Println("the student behind kyle is:", kyle.Next)
  
}