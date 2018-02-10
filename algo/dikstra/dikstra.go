package main

import(
  "fmt"
  "os"
  "io/ioutil"
  "strings"
  "strconv"
)


type Node struct {
  city string
  edges []*Edge
  searched bool
  parent *Node
}

type Edge struct {
  destination string
  weight int
}

type Graph struct {
  graph map[string]*Node
}

func (n *Node) addEdge(e *Edge) {
  n.edges = append(n.edges, e)
}

func stringConv(s string) int {
  new_int, _ := strconv.Atoi(s)
  return new_int
}


func main() { 
  if len(os.Args) < 2 {
    fmt.Println("there is no file")
    return
  }

  filename := os.Args[1]

  data, _ := ioutil.ReadFile(filename)
  var g *Graph = new(Graph)
  g.graph = make(map[string]*Node)
  for _, line := range strings.Split(string(data), "\n") {
    if line == "" {
      continue
    }
    td := strings.Split(string(line), " ")
    var n *Node = new(Node)
    var e *Edge = new(Edge)
    n.city = td[0]
    e.destination = td[1]
    weight := stringConv(td[2])
    e.weight = weight
    n.addEdge(e)

    _, ok := g.graph[n.city]

    if !ok {
      g.graph[n.city] = n 
    }
  }
  x := g.graph["Memphis"].edges
  fmt.Println(x[0])
}