package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func getElves(in []int) []int {
	sums := []int{}
	curr := 0
	for _, line := range in {
		if line == 0 {
			sums = append(sums, curr)
			curr = 0
		} else {
			curr += line
		}
	}
	return sums
}

func part1(in []int) {
	sums := getElves(in)
	fmt.Printf("Biggest: %d\n", intmath.Max(sums...))
}

func part2(in []int) {
	sums := getElves(in)
	sort.Ints(sums)
	fmt.Printf("Last three: %d\n", intmath.Sum(sums[len(sums)-3:]...))
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

func input() []int {
	ret := []int{}
	for _, line := range fileinput.ReadLines("input.txt") {
		i, err := strconv.Atoi(line)
		if err != nil {
			ret = append(ret, 0)
		} else {
			ret = append(ret, i)
		}
	}
	return ret
}
