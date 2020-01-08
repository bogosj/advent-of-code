package main

import (
	"fmt"
	"github.com/bogosj/advent-of-code/intmath"
	"math/bits"
	"time"
)

func isWall(x, y int) bool {
	v := x*x + 3*x + 2*x*y + y + y*y + 1350
	return bits.OnesCount64(uint64(v))%2 == 1
}

type state struct {
	p     intmath.Point
	steps int
}

func (s state) nextStates() (ret []state) {
	for _, n := range s.p.Neighbors() {
		if n.X >= 0 && n.Y >= 0 && !isWall(n.X, n.Y) {
			ret = append(ret, state{p: n, steps: s.steps + 1})
		}
	}
	return
}

func minDistanceToPoint(x, y int) int {
	states := []state{{p: intmath.Point{X: 1, Y: 1}}}
	visited := map[intmath.Point]bool{}
	for len(states) > 0 {
		state := states[0]
		states = states[1:]
		if visited[state.p] {
			continue
		}
		visited[state.p] = true
		if state.p.X == x && state.p.Y == y {
			return state.steps
		}
		ns := state.nextStates()
		states = append(states, ns...)
	}
	return -1
}

func part1() {
	fmt.Println("Minimum distance to 31,39:", minDistanceToPoint(31, 39))
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
