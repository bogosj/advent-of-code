package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type point = intmath.Point

func input() (ret [][]string) {
	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, strings.Fields(line))
	}
	return
}

func getPoint(s string) point {
	f := strings.Split(s, ",")
	f[1] = strings.ReplaceAll(f[1], ":", "")
	return point{X: intmath.Atoi(f[0]), Y: intmath.Atoi(f[1])}
}

func getDist(s string) (x, y int) {
	f := strings.Split(s, "x")
	return intmath.Atoi(f[0]), intmath.Atoi(f[1])
}

func makeOverlapMap() map[point]int {
	ret := map[point]int{}
	for _, claim := range input() {
		p := getPoint(claim[2])
		x1, y1 := getDist(claim[3])
		for x := p.X; x < p.X+x1; x++ {
			for y := p.Y; y < p.Y+y1; y++ {
				ret[point{X: x, Y: y}]++
			}
		}
	}
	return ret
}

func findUniqueClaim() string {
	m := makeOverlapMap()
CLAIMS:
	for _, claim := range input() {
		p := getPoint(claim[2])
		x1, y1 := getDist(claim[3])
		for x := p.X; x < p.X+x1; x++ {
			for y := p.Y; y < p.Y+y1; y++ {
				if m[point{X: x, Y: y}] > 1 {
					continue CLAIMS
				}
			}
		}
		return claim[0]
	}
	return ""
}

func part1() {
	var count int
	for _, v := range makeOverlapMap() {
		if v > 1 {
			count++
		}
	}
	fmt.Println("The number of spots that are multi-claimed is:", count)
}

func part2() {
	fmt.Println("The unqiue claim ID is:", findUniqueClaim())
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
