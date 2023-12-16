package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/day21/springdroid"
)

func part1() {
	d := springdroid.New("input.txt")
	d.RunProgram("part1.txt")
}

func part2() {
	d := springdroid.New("input.txt")
	d.RunProgram("part2.txt")
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}
