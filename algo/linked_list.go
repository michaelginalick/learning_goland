package main


import(
  "fmt"
)

type student struct{
  age int
  weight int
  name string
	next *student
}

func main() {

	mike := &student{33, 200, "mike", nil}
  studentList := addNode(mike)
	greg := &student{34, 443, "greg", nil}
  studentList = addNodeBeginning(greg, studentList)

	for s := studentList; s != nil; s = s.next {
	  fmt.Println(s.age, s.weight, s.name)
	}
}


func addNode(newStudent *student) *student {
	return newStudent
}

func addNodeBeginning(newStudent, studentList *student) *student {
  newStudent.next = studentList
	return newStudent
}
