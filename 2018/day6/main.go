package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type point = intmath.Point

func input() (ret []point) {
	for _, line := range fileinput.ReadLines("input.txt") {
		f := strings.FieldsFunc(line, func(r rune) bool { return r == ',' || r == ' ' })
		ret = append(ret, point{X: intmath.Atoi(f[0]), Y: intmath.Atoi(f[1])})
	}
	return
}

func max(ps []point) (x, y int) {
	for _, p := range ps {
		x = intmath.Max(p.X, x)
		y = intmath.Max(p.Y, y)
	}
	return
}

func closestPoint(p point, ps []point) (ret int) {
	var pairs [][]int
	for i, op := range ps {
		d := p.ManhattanDistanceTo(op)
		pairs = append(pairs, []int{i, d})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	if pairs[0][1] == pairs[1][1] {
		return -1
	}
	return pairs[0][0]
}

func makeMap(ps []point) (ret map[point]int, edges map[int]bool) {
	ret = map[point]int{}
	edges = map[int]bool{}
	maxX, maxY := max(ps)
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			p := point{X: x, Y: y}
			cp := closestPoint(p, ps)
			ret[p] = cp
			if x == 0 || y == 0 || x == maxX || y == maxY {
				edges[cp] = true
			}
		}
	}
	return
}

func part1() {
	m, edges := makeMap(input())
	pointCount := map[int]int{}
	for _, v := range m {
		pointCount[v]++
	}
	var pairs [][]int
	for k, v := range pointCount {
		pairs = append(pairs, []int{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] > pairs[j][1] })
	fmt.Println(pairs)
	for _, v := range pairs {
		if edges[v[0]] {
			continue
		}
		fmt.Printf("The %dth point has the most coverage: %d\n", v[0], v[1])
		return
	}
}

func closeEnough(p point, points []point) bool {
	var dist int
	for _, op := range points {
		dist += p.ManhattanDistanceTo(op)
		if dist > 10000 {
			return false
		}
	}
	return true
}

func part2() {
	points := input()
	m, _ := makeMap(points)
	var closePoints int
	for k := range m {
		if closeEnough(k, points) {
			closePoints++
		}
	}
	fmt.Printf("There are %d points close enough\n", closePoints)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
