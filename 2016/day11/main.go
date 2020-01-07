package main

import (
	"fmt"
	"time"
)

type generators struct {
	e int
	g map[int][]string
}

func newGen() *generators {
	g := generators{}
	g.g = map[int][]string{
		1: []string{"TG", "TM", "LG", "SG"},
		2: []string{"LM", "SM"},
		3: []string{"PG", "PM", "RG", "RM"},
		4: nil,
	}
	g.e = 1
	return &g
}

func part1() {
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
