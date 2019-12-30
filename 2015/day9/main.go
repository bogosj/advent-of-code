package main

import (
	"fmt"
	"github.com/bogosj/advent-of-code/2015/day9/locations"
	"time"
)

func part1() {
	l := locations.New()
	l.Load("input.txt")
	p, d := l.ShortestPath()
	fmt.Printf("Shortest path is: %v = %d\n", p, d)
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
