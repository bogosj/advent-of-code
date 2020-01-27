package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

const (
	open   = '.'
	trees  = '|'
	lumber = '#'
)

type point = intmath.Point

func input() map[point]rune {
	ret := map[point]rune{}
	for y, line := range fileinput.ReadLines("input.txt") {
		for x, c := range line {
			ret[point{X: x, Y: y}] = c
		}
	}
	return ret
}

func neighbors(p point, m map[point]rune) (cTrees, cLumber int) {
	for _, v := range p.AllNeighbors() {
		switch m[v] {
		case trees:
			cTrees++
		case lumber:
			cLumber++
		}
	}
	return
}

func advance(m map[point]rune) map[point]rune {
	ret := map[point]rune{}
	for p, v := range m {
		cTrees, cLumber := neighbors(p, m)
		switch v {
		case open:
			ret[p] = open
			if cTrees >= 3 {
				ret[p] = trees
			}
		case trees:
			ret[p] = trees
			if cLumber >= 3 {
				ret[p] = lumber
			}
		case lumber:
			ret[p] = open
			if cLumber > 0 && cTrees > 0 {
				ret[p] = lumber
			}
		}
	}
	return ret
}

func resourceValue(m map[point]rune) int {
	var cTrees, cLumber int
	for _, v := range m {
		switch v {
		case trees:
			cTrees++
		case lumber:
			cLumber++
		}
	}
	return cTrees * cLumber
}

func part1() {
	state := input()
	for i := 0; i < 10; i++ {
		state = advance(state)
	}
	fmt.Printf("The resource value after 10 minutes is: %d\n", resourceValue(state))
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
