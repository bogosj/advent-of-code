package main

import (
	"fmt"
	"github.com/bogosj/advent-of-code/2015/day18/yardgif"
	"time"
)

func part1() {
	g := yardgif.New("input.txt")
	g.Animate(100)
	fmt.Println("Number of lights on after 100 steps:", g.LightsOn())
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
