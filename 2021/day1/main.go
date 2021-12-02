package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func part1(in []string) {
	i := 0
	count := 0
	for _, s := range in {
		if intmath.Atoi(s) > i {
			count++
		}
		i = intmath.Atoi(s)
	}
	fmt.Println("Part 1 answer:", count-1)
}

func part2(in []string) {
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	in = input()
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []string {
	ret := []string{}

	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, line)
	}

	return ret
}
