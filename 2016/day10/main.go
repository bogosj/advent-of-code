package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2016/day10/balancebots"
	"github.com/bogosj/advent-of-code/fileinput"
)

func input() []string {
	lines := fileinput.ReadLines("input.txt")
	return lines
}

func part1() {
	f := balancebots.New()
	f.LookFor = []int{17, 61}
	fmt.Println("Bot holding 17, 61:", f.RunInstructions(input()))
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
