package game

import "github.com/bogosj/advent-of-code/2019/computer"

type point struct {
	x, y int
}

type Game struct {
	c     *computer.Computer
	grid  map[point]int
	Score int
}

func New(c *computer.Computer) *Game {
	g := Game{c: c}
	return &g
}

func (g *Game) Hack() {
	g.c.Hack(0, 2)
}

func (g *Game) LoadGrid() int {
	g.grid = map[point]int{}
	for {
		x := g.c.Compute(0)
		y := g.c.Compute(0)
		t := g.c.Compute(0)

		g.grid[point{x, y}] = t

		if g.c.Halted {
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

func (g *Game) PlayGame() {
	g.grid = map[point]int{}
	for {
		j := g.joyMove()
		x := g.c.Compute(j)
		y := g.c.Compute(j)
		t := g.c.Compute(j)

		if x == -1 && y == 0 {
			g.Score = t
		} else {
			g.grid[point{x, y}] = t
		}

		if g.c.Halted {
			break
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
