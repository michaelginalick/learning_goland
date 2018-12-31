package main

import (
	g "./game"
	"fmt"
)

func main() {
	g := g.New()
	g.RollMany(20, 10)
	fmt.Println(g.Score())
}
