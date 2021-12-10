package main

import (
	"fmt"
	"sort"
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

func basinSize(p intmath.Point, in map[intmath.Point]int) (size int) {
	visited := map[intmath.Point]bool{}
	visited[p] = true
	size++
	toVisit := p.Neighbors()
	for len(toVisit) > 0 {
		np := toVisit[0]
		toVisit = toVisit[1:]
		val, ok := in[np]
		if ok && val != 9 && !visited[np] {
			visited[np] = true
			size++
			toVisit = append(toVisit, np.Neighbors()...)
		}
	}
	return
}

func part1(in map[intmath.Point]int) {
	risk := 0
	for _, p := range lowPoints(in) {
		risk += in[p] + 1
	}
	fmt.Println("Part 1 answer:", risk)
}

func part2(in map[intmath.Point]int) {
	sizes := []int{}
	for _, p := range lowPoints(in) {
		sizes = append(sizes, basinSize(p, in))
	}
	sort.Ints(sizes)
	fmt.Println("Part 2 answer:", intmath.Product(sizes[len(sizes)-3:]...))
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
