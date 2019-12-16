package main

import (
	"fmt"
	"time"

	"jamesbogosian.com/advent-of-code/2019/day15/droid"
)

func part1() {
	d := droid.New("input.txt")
	r := d.Walk()
	fmt.Println("Length of path to oxygen:", len(r))
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 complete in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 complete in", time.Since(start))
}
