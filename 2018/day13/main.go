package main

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type point = intmath.Point
type world = map[point]rune

func input() world {
	ret := world{}
	for y, line := range fileinput.ReadLinesRaw("input.txt") {
		for x, c := range line {
			ret[point{X: x, Y: y}] = c
		}
	}
	return ret
}

const (
	north = iota
	east
	south
	west
)

const (
	left = iota
	straight
	right
)

type cart struct {
	loc        point
	dir        int
	turnChoice int
}

func (c *cart) String() string {
	return fmt.Sprintf("%v, %v", c.loc, c.dir)
}

func (c *cart) advance(w world) {
	switch c.dir {
	case north:
		c.loc.Y--
	case south:
		c.loc.Y++
	case east:
		c.loc.X++
	case west:
		c.loc.X--
	}
	switch w[c.loc] {
	case '/':
		switch c.dir {
		case north:
			c.dir = east
		case south:
			c.dir = west
		case east:
			c.dir = north
		case west:
			c.dir = south
		}
	case '\\':
		switch c.dir {
		case north:
			c.dir = west
		case south:
			c.dir = east
		case east:
			c.dir = south
		case west:
			c.dir = north
		}
	case '+':
		switch c.turnChoice {
		case left:
			c.turnChoice = straight
			c.dir = (c.dir + 3) % 4
		case straight:
			c.turnChoice = right
		case right:
			c.turnChoice = left
			c.dir = (c.dir + 1) % 4
		}
	}
}

func cartsStart(w world) (ret []*cart) {
	for p, c := range w {
		nc := &cart{loc: p}
		switch c {
		case 'v':
			nc.dir = south
		case '^':
			nc.dir = north
		case '<':
			nc.dir = west
		case '>':
			nc.dir = east
		default:
			nc = nil
		}
		if nc != nil {
			ret = append(ret, nc)
		}
	}
	return
}

func sortCarts(carts []*cart) {
	sort.Slice(carts, func(i, j int) bool {
		if carts[i].loc.Y == carts[j].loc.Y {
			return carts[i].loc.X < carts[j].loc.X
		}
		return carts[i].loc.Y < carts[j].loc.Y
	})
}

func crashExists(carts []*cart) (point, error) {
	seen := map[point]bool{}
	for _, c := range carts {
		if seen[c.loc] {
			return c.loc, nil
		}
		seen[c.loc] = true
	}
	return point{}, errors.New("no crash detected")
}

func crashPoint() point {
	w := input()
	carts := cartsStart(w)
	sortCarts(carts)
	for {
		for _, c := range carts {
			c.advance(w)
			if p, err := crashExists(carts); err == nil {
				return p
			}
		}
		sortCarts(carts)
	}
}

func part1() {
	fmt.Println("The first crash point is:", crashPoint())
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
