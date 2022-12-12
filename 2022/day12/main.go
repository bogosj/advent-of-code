package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type elevationMap struct {
	e          map[intmath.Point]int
	start, end intmath.Point
}

type state struct {
	p     intmath.Point
	steps int
}

func (s state) nextStates(e elevationMap) (ret []state) {
	for _, n := range s.p.Neighbors() {
		delta := e.e[n] - e.e[s.p]
		if delta <= 1 {
			_, ok := e.e[n]
			if ok {
				ret = append(ret, state{p: n, steps: s.steps + 1})
			}
		}
	}
	return
}

func part1(in elevationMap) {
	states := []state{{p: in.start}}
	visited := map[intmath.Point]bool{}
	steps := 0
	count := 0
	for len(states) > 0 {
		state := states[0]
		states = states[1:]
		if visited[state.p] {
			continue
		}
		visited[state.p] = true
		if state.p == in.end {
			steps = state.steps
			break
		}
		ns := state.nextStates(in)
		states = append(states, ns...)
		count++
	}
	fmt.Printf("It takes %d steps to reach the end.\n", steps)
}

func part2(in elevationMap) {
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

func input() elevationMap {
	e := elevationMap{}
	e.e = map[intmath.Point]int{}
	for y, line := range fileinput.ReadLines("input.txt") {
		for x, c := range line {
			curr := intmath.Point{X: x, Y: y}
			e.e[curr] = int(c - 'a')
			if c == 'S' {
				e.e[curr] = 0
				e.start = curr
			} else if c == 'E' {
				e.e[curr] = 25
				e.end = curr
			}
		}
	}
	return e
}
