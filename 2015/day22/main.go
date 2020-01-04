package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2015/day22/wizardsim"
)

func part1() {
	fmt.Println("The minimum mana required is:", wizardsim.Simulate(false))
}

func part2() {
	fmt.Println("The minimum mana required on hard mode is:", wizardsim.Simulate(true))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
