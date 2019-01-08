package game

// Game game object
type Game struct {
	rolls []int
	currentRole int
}


//New gives a new game struct
func New() *Game {
	return &Game{rolls: make([]int, 21), currentRole: 0}
}

// Roll counts the current role pins
func (g *Game) Roll(pins int) {
	g.rolls[g.currentRole] = pins
	g.currentRole++
}

// RollMany rolls n times
func (g *Game) RollMany(n, pins int) {
	for i := 0; i < n; i++ {
		g.rolls[i] = pins
	}
}

//Score scores the game
func (g *Game) Score() int {
	
	gameScore := 0
	frameIndex := 0

	for frameIndex < 10 {
		if g.strike(frameIndex) {
			gameScore += 10
			gameScore += g.rolls[frameIndex+1]+g.rolls[frameIndex+2]
			frameIndex++
		} else if g.spare(frameIndex) {
			gameScore += 10
			gameScore += g.rolls[frameIndex+2]
			frameIndex+=2
		} else {
			g.calculateRegularFrame(gameScore, frameIndex)
			frameIndex+=2
		}
	}
	return gameScore
}

func (g *Game) spare(frameIndex int) bool {
	return g.rolls[frameIndex] + g.rolls[frameIndex+1] == 10
}

func (g *Game) strike(frameIndex int) bool {
	return g.rolls[frameIndex] == 10
}

func (g *Game) calculateRegularFrame(gameScore, frameIndex int) int {
	gameScore += g.rolls[frameIndex] + g.rolls[frameIndex+1]
	return gameScore
}
