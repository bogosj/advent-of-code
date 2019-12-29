package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func part1() {
	lines := fileinput.ReadLines("input.txt")
	i := 0
	for _, r := range lines[0] {
		switch r {
		case '(':
			i++
		case ')':
			i--
		}
	}
	fmt.Println("Floor:", i)
}

func part2() {
	lines := fileinput.ReadLines("input.txt")
	i := 0
	for j, r := range lines[0] {
		switch r {
		case '(':
			i++
		case ')':
			i--
		}
		if i < 0 {
			i = j + 1
			break
		}
	}
	fmt.Println("Character:", i)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
