package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2015/day14/reindeerolympics"
)

func part1() {
	race := reindeerolympics.New("input.txt")
	race.Run(2503)
	fmt.Println(race.Winner())
}

func part2() {
	race := reindeerolympics.New("input.txt")
	race.Run(2503)
	fmt.Println(race.Winner2())
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
