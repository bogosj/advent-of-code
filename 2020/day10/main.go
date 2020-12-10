package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func part1(in []int) (ones, threes int) {
	sort.Ints(in)
	ones = 1
	threes = 1
	for i := 0; i < len(in)-1; i++ {
		switch in[i+1] - in[i] {
		case 1:
			ones++
		case 3:
			threes++
		}
	}
	fmt.Printf("The ones * threes = %v\n", ones*threes)
	return
}

func part2(in []int) {
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []int {
	ret := []int{}

	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, intmath.Atoi(line))
	}

	return ret
}
