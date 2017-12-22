package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
  "strconv"
)


type student struct{
  name string
  age int
  ssn string
  next *student
}


func addBeginning(studentList, newStudent *student) *student {
    if studentList == nil {
				return studentList
		}

	  newStudent.next = studentList
    //add this to break out of loop
    studentList.next.next.next = nil
    return newStudent
}

func deleteList(studentList *student) *student {

		if studentList == nil {
				return studentList
		}


		for s := studentList; s != nil; s = s.next {
			delete(studentList, s)
		}

		return studentList
}




func addEnd(studentList, newStudent *student) *student {

  if studentList == nil {
    return studentList
  }

  for s := studentList; s != nil; s = s.next {
		if s.next == nil {
			s.next = newStudent
			break
		}
	}
  return studentList
}

func reverse(curr *student) *student{

  if curr.next == nil{
    return curr
  } else {
    newHead := reverse(curr.next)
    curr.next.next = curr
    curr.next = nil
    return newHead
  }
}


func findElementByName(curr string, studentList *student)  string {

  for s := studentList; s != nil; s = s.next {
    if s.name == curr {
      return s.name
      break
    }
  }
  return ""
}

func main() {

  var students *student
  students = nil

  if len(os.Args) < 2 {
    fmt.Println("there is no file")
    return
  }

  filename := os.Args[1]

  data, _ := ioutil.ReadFile(filename)

  for _, line := range strings.Split(string(data), "\n") {

    if line == "" {
      continue
    }

    td := strings.Split(string(line), " ")

    ts := &student {
      name: td[0],
      ssn:  td[2],
      next: students,
    }
    ts.age, _ = strconv.Atoi(td[1])
    students = ts
  }


  fmt.Println("Original----------")
  for s := students; s != nil; s = s.next {
    fmt.Println(s.name, s.age, s.ssn)
  }

  students = reverse(students)

  fmt.Println("\nReverse----------")
  for s := students; s != nil; s = s.next {
    fmt.Println(s.name, s.age, s.ssn)
  }


  newStudent := &student{
    name: "Mike",
    ssn: "332-334-445",
    age: 33,
    next: nil,
  }
  fmt.Println("\nAdd New Element at the end-------")
  students = addEnd(students, newStudent)
  for s := students; s != nil; s = s.next {
    fmt.Println(s.name, s.age, s.ssn)
  }

	fmt.Println("\nAdd New Element at the beginning------")
  students = addBeginning(students, newStudent)
	for s := students; s != nil; s = s.next {
	  fmt.Println(s.name, s.age, s.ssn)
	}

  fmt.Println("\nFind element in the list------")
  foundStudent := findElementByName("Mike", students)
  fmt.Println(foundStudent)

	fmt.Println("\n Delete the linked list")
	shouldBeNil := deleteList(students)
	fmt.Println(shouldBeNil)

}
