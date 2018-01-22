package main


import (
    "fmt"
    "os"
    "io/ioutil"
    // "strings"
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


func main() {
  var file File

  jsonFile, err := os.Open("bacon.json")
  
  var g *Graph = new(Graph)
  
  if err != nil {
    fmt.Println(err)
  }
  defer jsonFile.Close()
  byteValue, _ := ioutil.ReadAll(jsonFile)
  json.Unmarshal(byteValue, &file)


  for i := 0; i < len(file.Movies); i++ {
    var n *Node = new(Node)
    n.value = file.Movies[i].Title
    g.addNode(n)
    
    for j := 0; j < len(file.Movies[i].Cast); j++ {
      var n *Node = new(Node)
      n.value = file.Movies[i].Cast[j]
      g.addNode(n)
    }
  }
  fmt.Println(g.graph)
}