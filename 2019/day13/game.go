package main

type point struct {
	x, y int
}

type game struct {
	c     *computer
	grid  map[point]int
	score int
}

func (g *game) hack() {
	g.c.prog[0] = 2
}

func (g *game) loadGrid() {
	g.grid = map[point]int{}
	for {
		x, err := g.c.compute(0)
		y, err := g.c.compute(0)
		t, err := g.c.compute(0)

		g.grid[point{x, y}] = t

		if err != nil {
			break
		}
	}
}

func (g *game) itemX(i int) int {
	for p, v := range g.grid {
		if v == i {
			return p.x
		}
	}
	return 0
}

func (g *game) joyMove() int {
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

func (g *game) playGame() {
	g.grid = map[point]int{}
	for {
		j := g.joyMove()
		x, err := g.c.compute(j)
		y, err := g.c.compute(j)
		t, err := g.c.compute(j)

		if x == -1 && y == 0 {
			g.score = t
		} else {
			g.grid[point{x, y}] = t
		}

		if err != nil {
			break
		}
	}
}

func (g *game) printScreen() {

}

func (g *game) blockCount() (ret int) {
	for _, v := range g.grid {
		if v == 2 {
			ret++
		}
	}
	return
}
