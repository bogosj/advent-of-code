package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/day20/donutmaze"
)

func test() {
	m := donutmaze.New("basic.txt")
	fmt.Println("Basic min distance:", m.ShortestPath())
	m = donutmaze.New("large.txt")
	fmt.Println("Large min distance:", m.ShortestPath())
}

func part1() {
	m := donutmaze.New("input.txt")
	fmt.Println("Min distance:", m.ShortestPath())
}

func part2() {
	m := donutmaze.New("input.txt")
	m.IsRecursive = true
	fmt.Println("Part 2 min distance:", m.ShortestPath())
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}
