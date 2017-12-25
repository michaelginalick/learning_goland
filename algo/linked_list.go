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
	james := &student{45, 554, "james", nil}
	marshall := &student{56, 123, "marshall", nil}
	allison := &student{54, 234, "allison", nil}
  studentList = addNodeBeginning(greg, studentList)
	studentList = addNodeBeginning(james, studentList)
	studentList = addNodeBeginning(marshall, studentList)
  studentList = addNodeEnd(allison, studentList)

  printList(studentList)

	fmt.Println("\nFind student-----------------")
	findStudent := findNodeByName("mike", studentList)
	fmt.Println(findStudent.name)
	fmt.Println("\nDelete student")
	studentList = deleteNodeByName("allison", studentList)
	printList(studentList)
	fmt.Println("\nreverse list recurrsive")
	studentList = reverseRecurrsive(studentList)
	printList(studentList)

	fmt.Println("\nreverse list iterative")
	studentList = reverseIterative(studentList)
	printList(studentList)

	billy := &student{34, 443, "billy", nil}
	studentList = insertAfter("allison", billy, studentList)
	printList(studentList)
	studentList = insertBefore("mike", billy, studentList)
	printList(studentList)
}


func printList(studentList *student) *student{
	for s := studentList; s != nil; s = s.next {
	  fmt.Println(s)
	}
	return studentList
}


func insertBefore(name string, newStudent, studentList *student) *student {

	if studentList == nil {
		return studentList
	}

	s := studentList

	for s != nil {
		if s.next.name == name {
			newStudent.next = s.next
			s.next = newStudent
			break
		}
		s = s.next
	}
	return studentList
}


func insertAfter(name string, newStudent, studentList *student) *student {
	if studentList == nil {
		return studentList
	}

	s := studentList

	for s != nil {
		nextNode := s.next
		if s.name == name {
			 s.next = newStudent
			 newStudent.next = nextNode
			 break
		}
		s = s.next
	}

	return studentList
}

func deleteNodeByName(name string, studentList *student) *student{
  if studentList == nil {
	  return studentList
	}
  s := studentList
  //delete head
  if s.name == name {
  	s = s.next
  	return s
  }
  // delete rest
	for s.next != nil {
	  if s.next.name == name {
		  s.next = s.next.next
			return studentList
		}
		s = s.next
	}
	return studentList
}

func reverseRecurrsive(studentList *student) *student {
	if studentList == nil {
		return studentList
	}

	s := studentList

	if s.next == nil {
		return s
	} else {
		newHead := reverseRecurrsive(s.next)
		s.next.next = s
		s.next = nil
		return newHead
	}

}


func reverseIterative(studentList *student) *student {
	if studentList == nil {
		return studentList
	}

	curr := studentList
	prev := &student{}
	next := &student{}

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	curr = prev
	return curr
}


func findNodeByName(name string, studentList *student) *student {
	if studentList == nil {
    return studentList
	}

	for s := studentList; s !=nil; s = s.next {
		if s.name == name {
			return s
			break
		}
	}
	return studentList
}

func addNodeEnd(newStudent, studentList *student) *student {
  if studentList == nil {
	  return studentList
	}

	for s := studentList; s != nil; s = s.next {
	  if s.next == nil {
		  s.next = newStudent
			return studentList
		}
	}

	return studentList
}


func addNode(newStudent *student) *student {
	return newStudent
}

func addNodeBeginning(newStudent, studentList *student) *student {
  newStudent.next = studentList
	return newStudent
}
