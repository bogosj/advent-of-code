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

func (f *forest) slopeCheck(right, down int) int {
	c := 0
	x := 0
	for y := 0; y < len(f.area); y += down {
		if f.isTree(x, y) {
			c++
		}
		x += right
	}
	return c
}

func part1(in forest) {
	c := in.slopeCheck(3, 1)
	fmt.Printf("Found %v trees along the path.\n", c)
}

func part2(in forest) {
	a := (in.slopeCheck(1, 1) *
		in.slopeCheck(3, 1) *
		in.slopeCheck(5, 1) *
		in.slopeCheck(7, 1) *
		in.slopeCheck(1, 2))
	fmt.Printf("Multipled slopes gets you %v.\n", a)
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
