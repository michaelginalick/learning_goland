package cell

// Cell that populates the board
type Cell struct {
	Revealed      bool
	X             int
	Y             int
	Bee           bool
	Display       string
	NeighborCount int
}

func (c *Cell) floodFill(board [][]*Cell) {
	for xOff := -1; xOff <= 1; xOff++ {
		i := c.X + xOff
		if i < 0 || i >= len(board) {
			continue
		}

		for yOff := -1; yOff <= 1; yOff++ {
			j := c.Y + yOff
			if j < 0 || j >= len(board) {
				continue
			}

			neighbor := board[i][j]

			if !neighbor.Revealed {
				neighbor.Show(board)
			}
		}
	}
}

// Show sets the cell to Revealed
func (c *Cell) Show(board [][]*Cell) {
	c.Revealed = true

	if c.NeighborCount == 0 {
		c.floodFill(board)
	}
}

// CountBeeNeighbors counts the number of bees neighboring the cell
func (c *Cell) CountBeeNeighbors(board [][]*Cell) {

	if c.Bee {
		c.NeighborCount = -1
		return
	}

	total := 0

	for xOff := -1; xOff <= 1; xOff++ {
		i := c.X + xOff
		if i < 0 || i >= len(board) {
			continue
		}

		for yOff := -1; yOff <= 1; yOff++ {
			j := c.Y + yOff
			if j < 0 || j >= len(board) {
				continue
			}

			neighbor := board[i][j]

			if neighbor.Bee {
				total++
			}
		}
	}
	c.NeighborCount = total
}
