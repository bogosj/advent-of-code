package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type point = intmath.Point

type world struct {
	w     map[point]bool
	units []*unit
}

func (w *world) sortUnits() {
	sort.Slice(w.units, func(i, j int) bool {
		if w.units[i].loc.Y == w.units[j].loc.Y {
			return w.units[i].loc.X < w.units[j].loc.X
		}
		return w.units[i].loc.Y < w.units[j].loc.Y
	})
}

func (w *world) isFreePoint(p point) bool {
	if !w.w[p] {
		return false
	}
	for _, u := range w.units {
		if u.hp > 0 && u.loc == p {
			return false
		}
	}
	return true
}

func (w *world) allOfTypeDead(r rune) bool {
	return w.hpOf(r) == 0
}

func (w *world) hpOf(r rune) (sum int) {
	for _, u := range w.units {
		if u.t == r && u.hp > 0 {
			sum += u.hp
		}
	}
	return
}

func (w *world) print() {
	for y := 0; y < 32; y++ {
	OUTER:
		for x := 0; x < 32; x++ {
			p := point{x, y}
			for _, u := range w.units {
				if u.loc == p && u.hp > 0 {
					fmt.Print(string(u.t))
					continue OUTER
				}
			}
			if w.w[p] {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		for _, u := range w.units {
			if u.loc.Y == y {
				fmt.Printf(" %s ", u)
			}
		}
		fmt.Println()
	}
}

func (w *world) simulate() {
	for i := 0; i < 1000; i++ {
		w.sortUnits()
		w.print()
		for _, u := range w.units {
			if u.hp <= 0 {
				continue
			}
			if w.allOfTypeDead('E') || w.allOfTypeDead('G') {
				fmt.Println("Done on round:", i)
				fmt.Println("Goblin HP:", w.hpOf('G'))
				fmt.Println("Elf HP:", w.hpOf('E'))
				return
			}
			u.takeTurn(w)
		}
		fmt.Printf("Turn %d done\n", i)
	}
}

func input() *world {
	w := world{}
	w.w = map[point]bool{}
	for y, line := range fileinput.ReadLines("input.txt") {
		for x, c := range line {
			p := point{X: x, Y: y}
			switch c {
			case '.':
				w.w[p] = true
			case 'G', 'E':
				w.w[p] = true
				w.units = append(w.units, &unit{hp: 200, atk: 3, loc: p, t: c})
			}
		}
	}
	return &w
}

type unit struct {
	hp, atk int
	loc     point
	t       rune
}

func (u *unit) String() string {
	return fmt.Sprintf("%s(%d)", string(u.t), u.hp)
}

func (u *unit) findTargets(w *world) (ret []*unit) {
	for _, ou := range w.units {
		if ou.t != u.t && ou.hp > 0 {
			ret = append(ret, ou)
		}
	}
	return
}

func (u *unit) attack(targets []*unit) bool {
	var c []*unit
	for _, n := range u.loc.Neighbors() {
		for _, t := range targets {
			if t.loc == n {
				c = append(c, t)
			}
		}
	}
	if len(c) == 0 {
		return false
	}
	sort.Slice(c, func(i, j int) bool {
		if c[i].hp == c[j].hp {
			if c[i].loc.Y == c[j].loc.Y {
				return c[i].loc.X < c[j].loc.X
			}
			return c[i].loc.Y < c[j].loc.Y
		}
		return c[i].hp < c[j].hp
	})
	c[0].hp -= u.atk
	return true
}

func (u *unit) spotsToMoveTo(w *world, targets []*unit) (ret []point) {
	for _, t := range targets {
		for _, n := range t.loc.Neighbors() {
			if w.w[n] {
				ret = append(ret, n)
			}
		}
	}
	return
}

type path struct {
	p     point
	moves string
	v     map[point]bool
}

func step(from, to point) string {
	if to.Y < from.Y {
		return "N"
	}
	if to.Y > from.Y {
		return "S"
	}
	if to.X > from.X {
		return "E"
	}
	return "W"
}

func (u *unit) pathsTo(p point, w *world) (ret []path) {
	var paths []path
	for _, n := range u.loc.Neighbors() {
		if w.isFreePoint(n) {
			p := path{p: n, moves: step(u.loc, n)}
			p.v = map[point]bool{u.loc: true}
			paths = append(paths, p)
		}
	}
	for len(paths) > 0 {
		curr := paths[0]
		paths = paths[1:]
		if curr.v[curr.p] {
			continue
		}
		if len(ret) > 0 {
			if len(curr.moves) > len(ret[0].moves) {
				continue
			}
		}
		if curr.p == p {
			ret = append(ret, curr)
		}
		for _, n := range curr.p.Neighbors() {
			if w.isFreePoint(n) {
				p := path{p: n, moves: curr.moves + step(curr.p, n)}
				p.v = curr.v
				p.v[curr.p] = true
				paths = append(paths, p)
			}
		}
	}
	return
}

func (u *unit) move(d rune) {
	switch d {
	case 'N':
		u.loc.Y--
	case 'W':
		u.loc.X--
	case 'E':
		u.loc.X++
	case 'S':
		u.loc.Y++
	}
}

func (u *unit) takeTurn(w *world) {
	// all opposing units
	targets := u.findTargets(w)
	if u.attack(targets) {
		return
	}
	// otherwise find all neighbor spots to all targets
	var paths []path
	for _, spot := range u.spotsToMoveTo(w, targets) {
		paths = append(paths, u.pathsTo(spot, w)...)
	}
	if len(paths) == 0 {
		return
	}
	sort.Slice(paths, func(i, j int) bool {
		l := map[byte]int{'N': 1, 'W': 2, 'E': 3, 'S': 4}
		pi, pj := paths[i], paths[j]
		if len(pi.moves) == len(pj.moves) {
			if pi.p == pj.p {
				return l[pi.moves[0]] < l[pj.moves[0]]
			}
			if pi.p.Y == pj.p.Y {
				return pi.p.X < pj.p.X
			}
			return pi.p.Y < pj.p.Y
		}
		return len(pi.moves) < len(pj.moves)
	})
	u.move(rune(paths[0].moves[0]))
	u.attack(targets)
}

func part1() {
	w := input()
	w.simulate()
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
