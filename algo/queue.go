package main

import(
	"fmt"
)

type Queue struct {
  slice int[]
}


func (q *Queue) Enqueue(i int) {
  q.slice = append(q.slice, i)
}

func (q *Queue) DeQueue() int {
  first := q.slice[0]

  q.slice = q.slice[1:len(q.slice)]
  return first
}



