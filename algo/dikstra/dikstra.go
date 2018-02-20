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
    from := td[0]
    to := td[1]
    weight := td[2]

    _, ok := g.graph[from]
    var fromV *Node = new(Node)
    if !ok {
      g.graph[from] = fromV
    } else {
      fromV = g.graph[from]
    }

    _, okTo := g.graph[to]
    var toV *Node = new(Node)
    if !okTo {
      g.graph[to] = toV
    } else {
      fromV = g.graph[to]
    }
    
    fromV.addEdge(toV, weight)


   
  }
  x := len(g.graph)
  fmt.Println(x)
}