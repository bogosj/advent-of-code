package main

import (
	"fmt"
	"time"

	"jamesbogosian.com/advent-of-code/2019/computer"
	"jamesbogosian.com/advent-of-code/2019/day11/robot"
)

func part1() {
	c := computer.New("input.txt")
	r := robot.New(c)
	r.Paint(0)
	panels := r.PrintHull()
	fmt.Println("Number of panels:", panels)
}

func part2() {
	c := computer.New("input.txt")
	r := robot.New(c)
	r.Paint(1)
	r.PrintHull()
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
