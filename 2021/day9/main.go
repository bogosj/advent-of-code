package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func lowPoints(in map[intmath.Point]int) []intmath.Point {
	ret := []intmath.Point{}
OUTER:
	for k, v := range in {
		for _, n := range k.Neighbors() {
			nv, ok := in[n]
			if nv <= v && ok {
				continue OUTER
			}
		}
		ret = append(ret, k)
	}
	return ret
}

func part1(in map[intmath.Point]int) {
	risk := 0
	points := lowPoints(in)
	for _, p := range points {
		risk += in[p] + 1
	}
	fmt.Println("Part 1 answer:", risk)
}

func part2(in map[intmath.Point]int) {
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
			ret[intmath.Point{X: x, Y: y}] = intmath.Atoi(string(char))
		}
	}

	return ret
}
