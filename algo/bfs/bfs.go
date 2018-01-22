package main


import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
)


type File struct {
  Movies []struct {
    Title string   `json:"title"`
    Cast  []string `json:"cast"`
  } `json:"movies"`
}

type Node struct {
  value string
  edges []*Node
  searched bool
  parent string
}

type Graph struct {
  nodes []*Node
  graph map[string]*Node
}

func (g *Graph) addNode(n *Node) {
  g.nodes = append(g.nodes, n)
  title := n.value

  if g.graph == nil {
    g.graph = make(map[string]*Node)
  }

  g.graph[title] = n
}


func (g *Graph) getNode(s string) *Node {
  for i := 0; i < len(g.nodes); i++ {
    if g.nodes[i].value == s {
      return g.nodes[i]
    }
  }
  var actorNode *Node = new(Node)
  return actorNode
}

func (n *Node) addEdge(actorNode *Node) {
  n.edges = append(n.edges, actorNode)
  actorNode.edges = append(actorNode.edges, n)
}


func main() {
  var file File
  var g *Graph = new(Graph)

  jsonFile, err := os.Open("bacon.json")
  
  if err != nil {
    fmt.Println(err)
  }

  defer jsonFile.Close()
  byteValue, _ := ioutil.ReadAll(jsonFile)
  json.Unmarshal(byteValue, &file)


  for i := 0; i < len(file.Movies); i++ {
    var movieNode *Node = new(Node)
    movieNode.value = file.Movies[i].Title
    g.addNode(movieNode)
    
    for j := 0; j < len(file.Movies[i].Cast); j++ {
      actorNode := g.getNode(file.Movies[i].Cast[j])

      actorNode.value = file.Movies[i].Cast[j]
      g.addNode(actorNode)
      movieNode.addEdge(actorNode)
      
    }
  }
  fmt.Println(g)
}