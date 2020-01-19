package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() (ret []int) {
	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, intmath.Atoi(line))
	}
	return
}

func part1() {
	fmt.Println("The frequency is:", intmath.Sum(input()...))
}

func part2() {
	seen := map[int]bool{}
	var f int
	for {
		for _, v := range input() {
			f += v
			if seen[f] {
				fmt.Println("The first frequency seen twice is:", f)
				return
			}
			seen[f] = true
		}
	}
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
