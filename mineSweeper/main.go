package main

import (
	"math/rand"
	"time"
	g "./game"
)

func main() {
	rand.Seed(time.Now().Unix())
	g := &g.Game{}
	g.StartGame()
}

