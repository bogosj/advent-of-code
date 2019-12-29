package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2015/day6/lights"
)

func part1() {
	l := lights.New()
	l.RunInstructions("input.txt")
	fmt.Println("Number of lights on:", l.LitLights())
}

func part2() {
	l := lights.New()
	l.RunInstructions2("input.txt")
	fmt.Println("Total brightness:", l.TotalBrightness())
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
