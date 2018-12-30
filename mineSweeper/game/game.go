package game

import (
	b "../board"
	"fmt"
	"bufio"
	"os"
	"strings"
)



// Game start the game
type Game struct {
	board b.Board
}

func (g *Game) createBoard() {
	g.board = b.Board{}
}

// StartGame start the game
func (g *Game) StartGame() {
	g.createBoard()
	b := g.board
	b.MakeGrid(15, 15)
	bCount := 20

	b.PopulateBoard()
	b.SeedBees(bCount)
	b.EstablishNeighbors()

	gameOver := false
	b.PrintBoard()

	for !gameOver {
		input := getCoords()
		cell := b.FindCell(input)
		if cell.Bee {
			gameOver = true
		}
		cell.Show(b.Grid)
		b.PrintBoard()
	}
	b.RevealAll()
	fmt.Println("GAME OVER")
	b.PrintBoard()
}

func getCoords() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter coordinates: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	coords := strings.Split(strings.TrimSpace(input), ",")
	return coords
}
