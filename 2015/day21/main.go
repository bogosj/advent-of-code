package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2015/day21/rpg"
)

func part1() {
	fmt.Println("The cheapest winning fighter costs:", rpg.CostOfCheapestWinningFighter())
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
