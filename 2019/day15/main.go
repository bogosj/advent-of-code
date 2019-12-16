package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/day15/droid"
	"github.com/bogosj/advent-of-code/2019/intmath"
)

func part1() {
	d := droid.New("input.txt")
	r := d.Walk()
	var lens []int
	for _, p := range r {
		lens = append(lens, len(p))
	}
	min := intmath.Min(lens...)
	fmt.Println("Length of path to oxygen:", min)
}

func part2() {
	d := droid.New("input.txt")
	d.Walk()
	m := d.ExpandO2()
	fmt.Println("Minutes to expand oxygen:", m)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 complete in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 complete in", time.Since(start))
}
