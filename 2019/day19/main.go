package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/computer"
	"github.com/bogosj/advent-of-code/2019/day19/beam"
)

func part1() {
	c := computer.New("input.txt")
	b := beam.New(c)
	fmt.Printf("Number of points: %v\n", b.Scan())
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
