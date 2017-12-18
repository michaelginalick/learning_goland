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

func main() {
  
  students := new(student)

  students.next = nil

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
      next: students.next,
    }
    ts.age, _ = strconv.Atoi(td[1])
    students.next = ts
  }

  

  for s := students.next; s != nil; s = s.next {
    fmt.Println(s.name, s.age, s.ssn)
  }

}