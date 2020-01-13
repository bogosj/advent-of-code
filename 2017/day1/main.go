package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func input() (ret []int) {
	lines := fileinput.ReadLines("input.txt")
	for _, c := range lines[0] {
		ret = append(ret, int(c-'0'))
	}
	return
}

func part1() {
	var sum int
	in := input()
	for i := 0; i < len(in); i++ {
		if in[i] == in[(i+1)%len(in)] {
			sum += in[i]
		}
	}
	fmt.Println("The solution is", sum)
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
