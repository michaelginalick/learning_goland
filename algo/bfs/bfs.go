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

type Queue struct {
  slice []*Node
}

func (q *Queue) EnQueue(n *Node) {
  q.slice = append(q.slice, n)
}

func (q *Queue) DeQueue() *Node {
  first := q.slice[0]
  q.slice = q.slice[1:len(q.slice)]

  return first
}

type Node struct {
  value string
  edges []*Node
  searched bool
  parent *Node
}

type Graph struct {
  nodes []*Node
  graph map[string]*Node
  end string
  start string
}

func (g *Graph) addNode(n *Node) {
  g.nodes = append(g.nodes, n)
  title := n.value

  if g.graph == nil {
    g.graph = make(map[string]*Node)
  }

  g.graph[title] = n
}


func (g *Graph) setStart(s string) *Node {
  for i := 0; i < len(g.nodes); i++ {
    if g.nodes[i].value == s {
      g.start = s
      return g.nodes[i]
    }    
  }
  var actorNode *Node = new(Node)
  return actorNode
}

func (g *Graph) setEnd(s string) *Node {
 for i := 0; i < len(g.nodes); i++ {
    if g.nodes[i].value == s {
      g.end = s
      return g.nodes[i]
    }    
  }
  var actorNode *Node = new(Node)
  return actorNode
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

func populateGraph(file File, g *Graph) {
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
}

func bfs(g *Graph) *Node {
  start := g.setStart("Rachel McAdams")
  end := g.setEnd("Kevin Bacon")
  var queue *Queue = new(Queue)

  start.searched = true
  queue.EnQueue(start)

  for len(queue.slice) > 0 {
    current := queue.DeQueue()
    if current.value == end.value {
      fmt.Println("found!")
      break
    }
    edges := current.edges

    for i := 0; i < len(edges); i++ {
      neighbor := edges[i]
      if !neighbor.searched {
        neighbor.searched = true
        neighbor.parent = current
        queue.EnQueue(neighbor)
      }
    }
  }

  return end
}

func buildPath(end *Node) {
  path := make([]*Node, 5)
  path = append(path, end)
  next := end.parent

  for n := next; n != nil; n = n.parent {
    path = append(path, n)
  }
  txt := ""
  for i:=len(path)-1; i >=0; i-- {
    l := path[i]
    if l != nil {
      txt += l.value + "-->"
    }
  } 
  fmt.Println(txt)
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

  populateGraph(file, g)
  end := bfs(g)
  buildPath(end)

}