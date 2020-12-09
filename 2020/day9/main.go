package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func part1(in []int) {
OUTER:
	for i := 25; i < len(in); i++ {
		for j := i - 25; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if in[j]+in[k] == in[i] {
					continue OUTER
				}
			}
		}
		fmt.Printf("The first bad number is %v\n", in[i])
		return
	}
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
