package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func charsUniqueIn(in string, count int) bool {
	s := map[rune]bool{}
	for _, r := range in {
		s[r] = true
	}
	return len(s) == count
}

func part1(in []string) {
	for i := 0; i < len(in[0]); i++ {
		if charsUniqueIn(in[0][i:i+4], 4) {
			fmt.Printf("%d characters need to be processed.\n", i+4)
			break
		}
	}
}

func part2(in []string) {
	for i := 0; i < len(in[0]); i++ {
		if charsUniqueIn(in[0][i:i+14], 14) {
			fmt.Printf("%d characters need to be processed.\n", i+14)
			break
		}
	}
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
	return fileinput.ReadLines("input.txt")
}
