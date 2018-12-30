package board

import (
	c "../cell"
	"math/rand"
	"strconv"
	"fmt"
)

// Board hold the grid for the game
type Board struct {
	Grid [][]*c.Cell
}


// PopulateBoard places a Cell object in each spot on the grid
func (b *Board) PopulateBoard() {
	for i := 0; i < len(b.Grid); i++ {
		for j := 0; j < len(b.Grid); j++ {
			b.Grid[i][j] = &c.Cell{Revealed: false, X: i, Y: j, Bee: false, Display: "#", NeighborCount: 0}
		}
	}
}

// SeedBees randomly places a number of Bees the board
func (b *Board) SeedBees(bCount int) {
	for bCount > 0 {
		n := rand.Intn(len(b.Grid))
		k := rand.Intn(len(b.Grid))
		randcell := b.Grid[n][k]
		if !randcell.Bee {
			randcell.Bee = true
		} else {
			continue
		}
		bCount--
	}
}

// EstablishNeighbors counts the neighboring Bees for each grid item
func (b *Board) EstablishNeighbors() {
	for i := 0; i < len(b.Grid); i++ {
		for j := 0; j < len(b.Grid); j++ {
			cell := b.Grid[i][j]
			cell.CountBeeNeighbors(b.Grid)
		}
	}
}

// FindCell returns the cell giving a pair of coordinates
func (b *Board) FindCell(coords []string) *c.Cell {
	xCord, _ := strconv.Atoi(coords[0])
	yCord, _ := strconv.Atoi(coords[1])
	cell := b.Grid[xCord][yCord]

	return cell
}

// PrintBoard Pretty Prints the board
func (b *Board) PrintBoard() {
	for x := 0; x < len(b.Grid); x++ {
		for y := 0; y < len(b.Grid); y++ {
			if b.Grid[x][y].Revealed {
				if b.Grid[x][y].Bee {
					fmt.Print("B")
				} else {
					if b.Grid[x][y].NeighborCount > 0 {
						fmt.Print(b.Grid[x][y].NeighborCount)
					} else {
						fmt.Print(" ")
					}
				}
			} else {
				fmt.Print(b.Grid[x][y].Display)
			}
			fmt.Print("\t")
		}
		fmt.Println()
	}
}

// MakeGrid make the 2D grid
func (b *Board) MakeGrid(x, y int) {
	twoD := make([][]*c.Cell, x)

	for i := 0; i < x; i++ {
		twoD[i] = make([]*c.Cell, y)
	}
	b.Grid = twoD
}

// RevealAll Sets every cell revealved to true
func (b *Board) RevealAll() {
	for i := 0; i < len(b.Grid); i++ {
		for j := 0; j < len(b.Grid); j++ {
			b.Grid[i][j].Revealed = true
		}
	}
}
