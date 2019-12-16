package droid

import (
	"fmt"

	"github.com/bogosj/advent-of-code/2019/computer"
	"github.com/bogosj/advent-of-code/2019/intmath"
)

type state struct {
	// Whether these directions have been explored.
	eN, eS, eW, eE bool
	isWall         bool
	isO2           bool
}

// Droid represents a repair droid.
type Droid struct {
	c   *computer.Computer
	m   map[intmath.Point]*state
	cur intmath.Point
	opp map[int]int
}

// New creates a new Droid based on the provided file path program.
func New(n string) *Droid {
	c := computer.New(n)
	d := Droid{c: c}
	d.m = map[intmath.Point]*state{}
	s := state{}
	d.m[intmath.Point{}] = &s
	d.opp = map[int]int{1: 2, 2: 1, 3: 4, 4: 3}
	return &d
}

// pointInDir returns the potential new point to explore, and the state of that point.
func pointInDir(p intmath.Point, dir int) (intmath.Point, *state) {
	switch dir {
	case 1:
		return intmath.Point{X: p.X, Y: p.Y + 1}, &state{eS: true}
	case 2:
		return intmath.Point{X: p.X, Y: p.Y - 1}, &state{eN: true}
	case 3:
		return intmath.Point{X: p.X - 1, Y: p.Y}, &state{eE: true}
	case 4:
		return intmath.Point{X: p.X + 1, Y: p.Y}, &state{eW: true}
	}
	return intmath.Point{}, nil
}

func (d *Droid) move(dir int) bool {
	p, ns := pointInDir(d.cur, dir)

	switch dir {
	case 1:
		d.m[d.cur].eN = true
	case 2:
		d.m[d.cur].eS = true
	case 3:
		d.m[d.cur].eW = true
	case 4:
		d.m[d.cur].eE = true
	}

	out := d.c.Compute(dir)
	if out == 0 {
		ns.isWall = true
		d.m[p] = ns
		return false
	}

	d.cur = p
	_, ok := d.m[d.cur]
	if !ok {
		d.m[d.cur] = ns
	}

	if out == 2 {
		d.m[d.cur].isO2 = true
	}

	return true
}

// Walk makes the droid walk and find all paths to the goal.
func (d *Droid) Walk() (ret [][]int) {
	path := []int{}
	for {
		s := d.m[d.cur]
		dir := 0
		if !s.eN {
			dir = 1
		} else if !s.eS {
			dir = 2
		} else if !s.eW {
			dir = 3
		} else if !s.eE {
			dir = 4
		}
		if dir == 0 {
			// Backtrack
			if len(path) == 0 {
				return
			}
			dir, path = path[len(path)-1], path[:len(path)-1]
			d.move(d.opp[dir])
		} else {
			moved := d.move(dir)
			if moved {
				path = append(path, dir)
			}
			if d.m[d.cur].isO2 {
				ret = append(ret, path)
			}
			continue
		}
	}
}

// PrintMap renders the map to stdout.
func (d *Droid) PrintMap() {
	var allX, allY []int
	for k := range d.m {
		allX = append(allX, k.X)
		allY = append(allY, k.Y)
	}
	minX := intmath.Min(allX...)
	minY := intmath.Min(allY...)
	maxX := intmath.Max(allX...)
	maxY := intmath.Max(allY...)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			s, ok := d.m[intmath.Point{X: x, Y: y}]
			if ok {
				if s.isWall {
					fmt.Print("#")
				} else if s.isO2 {
					fmt.Print("O")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print("?")
			}
		}
		fmt.Println()
	}
}

func (d *Droid) lackO2Points() (ret []intmath.Point) {
	for k, v := range d.m {
		if !v.isO2 && !v.isWall {
			ret = append(ret, k)
		}
	}
	return
}

func (d *Droid) o2Points() (ret []intmath.Point) {
	for k, v := range d.m {
		if v.isO2 {
			ret = append(ret, k)
		}
	}
	return
}

// ExpandO2 redraws the map placing oxygen in cells that neighbor other oxygen.
// It returns the number of minutes it took to fill the space with oxygen.
func (d *Droid) ExpandO2() (minutes int) {
	for len(d.lackO2Points()) > 0 {
		o2p := d.o2Points()
		for _, p := range o2p {
			ns := p.Neighbors()
			for _, n := range ns {
				if newO2Point, ok := d.m[n]; ok {
					if !newO2Point.isWall {
						newO2Point.isO2 = true
					}
				}
			}
		}
		minutes++
	}
	return
}
