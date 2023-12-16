package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/day25/cryostasis"
)

func part1() {
	c := cryostasis.New("input.txt")
	c.Run()
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}
