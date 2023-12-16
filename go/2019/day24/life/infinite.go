package life

import (
	"fmt"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

// Infinite represents an infinite series of boards.
type Infinite struct {
	world map[int][5][5]int
}

// NewInfinite creates a new infinite series of boards.
func NewInfinite(p string) *Infinite {
	i := Infinite{}
	i.world = map[int][5][5]int{}
	lines := fileinput.ReadLines(p)
	init := [5][5]int{}
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				init[y][x] = 1
			}
		}
	}
	i.world[0] = init
	return &i
}

// BugCount returns the number of bugs in this infinite space.
func (i *Infinite) BugCount() (ret int) {
	for _, l := range i.world {
		for _, row := range l {
			for _, c := range row {
				ret += c
			}
		}
	}
	return
}

func (i *Infinite) advance() {
	next := map[int][5][5]int{}
	for z := -200; z <= 200; z++ {
		nl := [5][5]int{}
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if x == 2 && y == 2 {
					continue
				}
				p := intmath.Point{X: x, Y: y}
				var bugCount int
				for _, n := range p.Neighbors() {
					switch {
					case n.X == -1:
						bugCount += i.world[z-1][2][1]
					case n.Y == -1:
						bugCount += i.world[z-1][1][2]
					case n.X == 5:
						bugCount += i.world[z-1][2][3]
					case n.Y == 5:
						bugCount += i.world[z-1][3][2]
					case n.X == 2 && n.Y == 2:
						var xs, ys []int
						if x == 2 {
							xs = []int{0, 1, 2, 3, 4}
							ys = []int{4}
							if y == 1 {
								ys[0] = 0
							}
						}
						if y == 2 {
							ys = []int{0, 1, 2, 3, 4}
							xs = []int{4}
							if x == 1 {
								xs[0] = 0
							}
						}
						for _, bx := range xs {
							for _, by := range ys {
								bugCount += i.world[z+1][by][bx]
							}
						}
					default:
						bugCount += i.world[z][n.Y][n.X]
					}
				}
				nl[y][x] = 0
				if bugCount == 1 {
					nl[y][x] = 1
				}
				if bugCount == 2 && i.world[z][y][x] == 0 {
					nl[y][x] = 1
				}
			}
		}
		next[z] = nl
	}
	i.world = next
}

// Run steps the infinite world ahead t minutes.
func (i *Infinite) Run(t int) {
	for j := 0; j < t; j++ {
		i.advance()
	}
}

func (i *Infinite) Print() {
	for z := -5; z <= 5; z++ {
		fmt.Println()
		fmt.Println("Depth", z)
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if x == 2 && y == 2 {
					fmt.Print("?")
				} else {
					if i.world[z][y][x] == 0 {
						fmt.Print(".")
					} else {
						fmt.Print("#")
					}
				}
			}
			fmt.Println()
		}
	}
}
