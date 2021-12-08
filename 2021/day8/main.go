package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func part1(in []string) {
	count := 0
	for _, line := range in {
		f := strings.Split(line, " | ")
		for _, digit := range strings.Fields(f[1]) {
			if len(digit) != 5 && len(digit) != 6 {
				count++
			}
		}
	}
	fmt.Println("Part 1 answer:", count)
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
