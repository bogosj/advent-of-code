package game

import "jamesbogosian.com/advent-of-code/2019/computer"

type point struct {
	x, y int
}

type Game struct {
	C     *computer.Computer
	grid  map[point]int
	Score int
}

func (g *Game) Hack() {
	g.C.Prog[0] = 2
}

func (g *Game) LoadGrid() int {
	g.grid = map[point]int{}
	for {
		x, err := g.C.Compute(0)
		y, err := g.C.Compute(0)
		t, err := g.C.Compute(0)

		g.grid[point{x, y}] = t

		if err != nil {
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
		x, err := g.C.Compute(j)
		y, err := g.C.Compute(j)
		t, err := g.C.Compute(j)

		if x == -1 && y == 0 {
			g.Score = t
		} else {
			g.grid[point{x, y}] = t
		}

		if err != nil {
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
