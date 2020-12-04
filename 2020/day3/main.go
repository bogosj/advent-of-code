package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

type forest struct {
	area []string
}

func (f *forest) isTree(x, y int) bool {
	row := f.area[y]
	return row[x%len(row)] == '#'
}

func part1(in forest) {
	c := 0
	x := 0
	for y := range in.area {
		if in.isTree(x, y) {
			c++
		}
		x += 3
	}
	fmt.Printf("Found %v trees along the path.\n", c)
}

func part2(in forest) {
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() forest {
	return forest{
		area: fileinput.ReadLines("input.txt"),
	}
}
