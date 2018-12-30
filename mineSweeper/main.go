package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	c "./cell"
)

func main() {
	rand.Seed(time.Now().Unix())
	board := makeGrid(15, 15)
	bCount := 20
	populateBoard(board)
	seedBees(board, bCount)
	establishNeighbors(board)
	gameOver := false
	printBoard(board)

	for !gameOver {
		input := getCoords()
		cell := findCell(input, board)
		if cell.Bee {
			gameOver = true
		}
		cell.Show(board)
		printBoard(board)
	}

	printBoard(board)
}

func seedBees(board [][]*c.Cell, bCount int) {
	for bCount > 0 {
		n := rand.Intn(len(board))
		k := rand.Intn(len(board))
		randcell := board[n][k]
		if !randcell.Bee {
			randcell.Bee = true
		} else {
			continue
		}
		bCount--
	}
}

func populateBoard(board [][]*c.Cell) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			board[i][j] = &c.Cell{Revealed: false, X: i, Y: j, Bee: false, Display: "#", NeighborCount: 0}
		}
	}
}

func findCell(coords []string, board [][]*c.Cell) *c.Cell {

	xCord, _ := strconv.Atoi(coords[0])
	yCord, _ := strconv.Atoi(coords[1])
	cell := board[xCord][yCord]

	return cell
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

func establishNeighbors(board [][]*c.Cell) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			cell := board[i][j]
			cell.CountNeighbors(board)
		}
	}
}

func printBoard(board [][]*c.Cell) {

	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board); y++ {
			if board[x][y].Revealed {
				if board[x][y].Bee {
					fmt.Print("B")
				} else {
					if board[x][y].NeighborCount > 0 {
						fmt.Print(board[x][y].NeighborCount)
					} else {
						fmt.Print(" ")
					}
				}
			} else {
				fmt.Print(board[x][y].Display)
			}
			fmt.Print("\t")
		}
		fmt.Println()
	}

}

func makeGrid(x, y int) [][]*c.Cell {

	twoD := make([][]*c.Cell, x)

	for i := 0; i < x; i++ {
		twoD[i] = make([]*c.Cell, y)
	}

	return twoD

}
