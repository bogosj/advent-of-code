package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
	pq "github.com/jupp0r/go-priority-queue"
)

type state struct {
	pos  intmath.Point
	risk int
}

func findExit(ceiling map[intmath.Point]int) intmath.Point {
	xs, ys := []int{}, []int{}
	for point := range ceiling {
		xs = append(xs, point.X)
		ys = append(ys, point.Y)
	}
	return intmath.Point{X: intmath.Max(xs...), Y: intmath.Max(ys...)}
}

func optimalPath(ceiling map[intmath.Point]int) *state {
	exit := findExit(ceiling)
	states := pq.New()
	states.Insert(&state{}, 0)
	visited := map[intmath.Point]bool{}
	visited[intmath.Point{}] = true

	for states.Len() > 0 {
		i, err := states.Pop()
		if err != nil {
			panic(err)
		}
		curr := i.(*state)
		if curr.pos.ManhattanDistanceTo(exit) == 0 {
			return curr
		}
		for _, n := range curr.pos.Neighbors() {
			if visited[n] {
				continue
			}
			risk, ok := ceiling[n]
			if ok {
				ns := &state{pos: n, risk: curr.risk + risk}
				states.Insert(ns, float64(ns.risk))
				visited[ns.pos] = true
			}
		}
	}
	return &state{}
}

func part1(ceiling map[intmath.Point]int) {
	s := optimalPath(ceiling)
	fmt.Println("Part 1 answer:", s.risk)
}

func part2(ceiling map[intmath.Point]int) {
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

func input() map[intmath.Point]int {
	ret := map[intmath.Point]int{}

	for y, line := range fileinput.ReadLines("input.txt") {
		for x, char := range line {
			p := intmath.Point{X: x, Y: y}
			ret[p] = intmath.Atoi(string(char))
		}
	}

	return ret
}
