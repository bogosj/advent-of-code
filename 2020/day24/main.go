package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

type hexPoint struct {
	x, y, z int
}

func (h *hexPoint) neighbors() []hexPoint {
	return []hexPoint{
		{h.x + 1, h.y - 1, h.z}, {h.x + 1, h.y, h.z - 1}, {h.x, h.y + 1, h.z - 1},
		{h.x - 1, h.y + 1, h.z}, {h.x - 1, h.y, h.z + 1}, {h.x, h.y - 1, h.z + 1},
	}
}

type directions struct {
	d string
}

func (d *directions) endPoint() (ret hexPoint) {
	for i := 0; i < len(d.d); i++ {
		switch d.d[i] {
		case 'e':
			ret.x++
			ret.y--
		case 'w':
			ret.x--
			ret.y++
		case 'n':
			i++
			ret.z--
			if d.d[i] == 'w' {
				ret.y++
			} else {
				ret.x++
			}
		case 's':
			i++
			ret.z++
			if d.d[i] == 'w' {
				ret.x--
			} else {
				ret.y--
			}
		}
	}
	return
}

func startingTiles(in []directions) map[hexPoint]bool {
	ret := map[hexPoint]bool{}
	for _, d := range in {
		p := d.endPoint()
		ret[p] = !ret[p]
	}
	return ret
}

func part1(in []directions) {
	m := startingTiles(in)
	count := 0
	for _, v := range m {
		if v {
			count++
		}
	}
	fmt.Printf("There are %d black tiles\n", count)
}

func countNeighbors(p hexPoint, state map[hexPoint]bool) (black int) {
	for _, n := range p.neighbors() {
		if state[n] {
			black++
		}
	}
	return
}

func step(state map[hexPoint]bool) map[hexPoint]bool {
	ret := map[hexPoint]bool{}

	// Flip blacks to white.
	for p, black := range state {
		if black {
			nBlack := countNeighbors(p, state)
			if nBlack == 0 || nBlack > 2 {
				ret[p] = false
			} else {
				ret[p] = true
			}
		}
	}

	// Find neighbor whites of exactly two blacks and flip them black.
	whitesWithBlack := map[hexPoint]int{}
	for p, black := range state {
		if black {
			for _, np := range p.neighbors() {
				if !state[np] {
					whitesWithBlack[np]++
				}
			}
		}
	}
	for p, c := range whitesWithBlack {
		if c == 2 {
			ret[p] = true
		}
	}
	return ret
}

func part2(in []directions) {
	m := startingTiles(in)
	for i := 0; i < 100; i++ {
		m = step(m)
	}
	c := 0
	for _, black := range m {
		if black {
			c++
		}
	}
	fmt.Printf("After 100 days there will be %v black tiles\n", c)
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	in = input()
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []directions {
	ret := []directions{}

	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, directions{d: line})
	}

	return ret
}
