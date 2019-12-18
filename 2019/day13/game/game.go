package game

import (
	"github.com/bogosj/advent-of-code/2019/computer"
)

type point struct {
	x, y int
}

// Game represents the breakout game cabinet.
type Game struct {
	c     *computer.Computer
	grid  map[point]int
	Score int
}

// New creates a new breakout game with the provided computer.
func New(c *computer.Computer) *Game {
	g := Game{c: c}
	return &g
}

// Hack sets the underlying computer to allow playing for free.
func (g *Game) Hack() {
	g.c.Hack(0, 2)
}

// LoadGrid reads the grid into memory.
func (g *Game) LoadGrid() int {
	g.grid = map[point]int{}
	in := make(chan int, 1)
	out := g.c.Compute(in)
	in <- 0
	for {
		x := <-out
		y := <-out
		t, ok := <-out

		g.grid[point{x, y}] = t

		if !ok {
			break
		}
	}
	return g.blockCount()
}

func (g *Game) itemX(i int) int {
	for p, v := range g.grid {
		if v == i {
			return p.x
		}
	}
	return 0
}

func (g *Game) joyMove() int {
	b := g.itemX(4)
	p := g.itemX(3)
	if b > p {
		return 1
	}
	if p > b {
		return -1
	}
	return 0
}

// PlayGame plays breakout.
func (g *Game) PlayGame() {
	g.grid = map[point]int{}
	in := make(chan int, 1)
	out := g.c.Compute(in)
	for {
		x, ok := <-out
		if !ok {
			break
		}
		y, z := <-out, <-out

		if x == -1 && y == 0 {
			g.Score = z
		} else {
			g.grid[point{x, y}] = z
		}
		if z == 4 {
			in <- g.joyMove()
		}
	}
}

func (g *Game) blockCount() (ret int) {
	for _, v := range g.grid {
		if v == 2 {
			ret++
		}
	}
	return
}
