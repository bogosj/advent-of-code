package yardgif

import "github.com/bogosj/advent-of-code/fileinput"

import "github.com/bogosj/advent-of-code/intmath"

const (
	on  = '#'
	off = '.'
)

// Grid represents a grid of lights
type Grid struct {
	g [][]rune
}

// New returns a new grid based on the provided file path.
func New(p string) *Grid {
	g := Grid{}
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		g.g = append(g.g, []rune(line))
	}
	return &g
}

// LightsOn returns the number of lights that are on.
func (g *Grid) LightsOn() (ret int) {
	for _, row := range g.g {
		for _, cell := range row {
			if cell == on {
				ret++
			}
		}
	}
	return
}

func (g *Grid) lightIsOn(p intmath.Point) bool {
	if p.X < 0 || p.Y < 0 || p.Y >= len(g.g) || p.X >= len(g.g[0]) {
		return false
	}
	return g.g[p.Y][p.X] == on
}

func (g *Grid) animate() {
	var ng [][]rune
	for y, row := range g.g {
		var nr []rune
		for x, cell := range row {
			i := 0
			p := intmath.Point{X: x, Y: y}
			for _, n := range p.AllNeighbors() {
				if g.lightIsOn(n) {
					i++
				}
			}
			if cell == off {
				if i == 3 {
					nr = append(nr, on)
				} else {
					nr = append(nr, off)
				}
			} else {
				if i == 2 || i == 3 {
					nr = append(nr, on)
				} else {
					nr = append(nr, off)
				}
			}
		}
		ng = append(ng, nr)
	}
	g.g = ng
}

// Animate runs the program for the provided number of steps.
func (g *Grid) Animate(steps int) {
	for i := 0; i < steps; i++ {
		g.animate()
	}
}
