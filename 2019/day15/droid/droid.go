package droid

import (
	"jamesbogosian.com/advent-of-code/2019/computer"
	"jamesbogosian.com/advent-of-code/2019/intmath"
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
func (d *Droid) Walk() (ret []int) {
	path := []int{}
	for {
		s := d.m[d.cur]
		dir := 0
		if !s.eN {
			dir = 1
		}
		if !s.eS {
			dir = 2
		}
		if !s.eW {
			dir = 3
		}
		if !s.eE {
			dir = 4
		}
		if dir == 0 {
			// Backtrack
			dir, path = path[len(path)-1], path[:len(path)-1]
			d.move(d.opp[dir])
		} else {
			moved := d.move(dir)
			if moved {
				path = append(path, dir)
			}
			if d.m[d.cur].isO2 {
				return path
			}
			continue
		}
	}
}
