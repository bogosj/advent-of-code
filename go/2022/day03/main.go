package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func makeCompartments(in []string) [][]string {
	ret := [][]string{}
	for _, line := range in {
		c := []string{}
		c = append(c, line[:len(line)/2])
		c = append(c, line[len(line)/2:])
		ret = append(ret, c)
	}
	return ret
}

func findMatchedRune(in []string) rune {
	runes := map[rune]bool{}
	for _, r := range in[0] {
		runes[r] = true
	}
	for _, r := range in[1] {
		if runes[r] {
			return r
		}
	}
	return '0'
}

func getPriority(r rune) int {
	if r >= 97 && r <= 122 {
		return int(r - 96)
	}
	if r >= 65 && r <= 90 {
		return int(r - 38)
	}
	return -1
}

func part1(in []string) {
	compartments := makeCompartments(in)
	sum := 0
	for _, pair := range compartments {
		sum += getPriority(findMatchedRune(pair))
	}
	fmt.Printf("Sum of priorities: %d\n", sum)
}

func getTripleRunes(in []string) []rune {
	ret := []rune{}
	for len(in) > 0 {
		runes := map[rune]int{}
		for _, r := range in[0] {
			runes[r] = 1
		}
		for _, r := range in[1] {
			if runes[r] == 1 {
				runes[r] = 2
			}
		}
		for _, r := range in[2] {
			if runes[r] == 2 {
				ret = append(ret, r)
				break
			}
		}
		in = in[3:]
	}
	return ret
}

func part2(in []string) {
	sum := 0
	for _, r := range getTripleRunes(in) {
		sum += getPriority(r)
	}
	fmt.Printf("Priority: %d\n", sum)
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
