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

func addEnd(studentList *student, newStudent *student) *student {

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
  fmt.Println("\nAdd New Element")
  students = addEnd(students, newStudent)
  for s := students; s != nil; s = s.next {
    fmt.Println(s.name, s.age, s.ssn)
  }

}