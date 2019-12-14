package main

import (
	"fmt"
	"time"

	"jamesbogosian.com/advent-of-code/2019/computer"
)

func diagnose() int {
	c := computer.New("input.txt")
	for {
		out := c.Compute(1)
		if out != 0 {
			return out
		}
	}
}

func part1() {
	out := diagnose()
	fmt.Println("Part 1:", out)
}

func part2() {
	c := computer.New("input.txt")
	out := c.Compute(5)
	fmt.Println("Part 2:", out)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
