package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type point = intmath.Point
type nodeState int
type direction int

const (
	up = iota
	right
	down
	left
)
const (
	clean = iota
	weakened
	infected
	flagged
)

type grid struct {
	g        map[point]nodeState
	virus    point
	virusDir direction
}

func newGrid() *grid {
	g := grid{g: map[point]nodeState{}}
	input := fileinput.ReadLines("input.txt")
	for y, line := range input {
		for x, c := range line {
			if c == '#' {
				g.g[point{X: x, Y: y}] = infected
			}
		}
	}
	g.virus = point{
		X: len(input[0]) / 2,
		Y: len(input) / 2,
	}
	g.virusDir = up
	return &g
}

func (g *grid) move() {
	switch g.virusDir {
	case up:
		g.virus.Y--
	case down:
		g.virus.Y++
	case left:
		g.virus.X--
	case right:
		g.virus.X++
	}
}

func (g *grid) step() (causedInfection bool) {
	if g.g[g.virus] == infected {
		g.turnRight()
		g.g[g.virus] = clean
	} else {
		g.turnLeft()
		g.g[g.virus] = infected
		causedInfection = true
	}
	g.move()
	return
}

func (g *grid) stepV2() (causedInfection bool) {
	switch g.g[g.virus] {
	case clean:
		g.turnLeft()
		g.g[g.virus] = weakened
	case weakened:
		g.g[g.virus] = infected
		causedInfection = true
	case infected:
		g.turnRight()
		g.g[g.virus] = flagged
	case flagged:
		g.turnRight()
		g.turnRight()
		g.g[g.virus] = clean
	}
	g.move()
	return
}

func (g *grid) turnRight() {
	g.virusDir = (g.virusDir + 1) % 4
}

func (g *grid) turnLeft() {
	g.virusDir = (g.virusDir + 3) % 4
}

func part1() {
	g := newGrid()
	var c int
	for i := 0; i < 10000; i++ {
		if g.step() == true {
			c++
		}
	}
	fmt.Printf("After 10000 burst of activity, %d caused infection \n", c)
}

func part2() {
	g := newGrid()
	var c int
	for i := 0; i < 10000000; i++ {
		if g.stepV2() == true {
			c++
		}
	}
	fmt.Printf("After 10000000 burst of activity with the new virus, %d caused infection \n", c)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
