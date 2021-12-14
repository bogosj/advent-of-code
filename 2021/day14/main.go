package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func processInput(in []string) (string, map[string]string) {
	ret := map[string]string{}
	for _, line := range in[2:] {
		f := strings.Split(line, " -> ")
		ret[f[0]] = f[1]
	}
	return in[0], ret
}

func stringToPairs(in string) map[string]int {
	ret := map[string]int{}
	for i := 0; i < len(in)-1; i++ {
		ret[in[i:i+2]]++
	}
	return ret
}

func step(pairs map[string]int, instructions map[string]string) map[string]int {
	ret := map[string]int{}
	for pair, count := range pairs {
		toAdd := instructions[pair]
		ret[string(pair[0])+toAdd] += count
		ret[toAdd+string(pair[1])] += count
	}
	return ret
}

func elementDiff(initial string, pairs map[string]int) int {
	elements := map[byte]int{}
	elements[initial[0]]++
	elements[initial[len(initial)-1]]++
	for pair, count := range pairs {
		elements[pair[0]] += count
		elements[pair[1]] += count
	}
	counts := []int{}
	for _, count := range elements {
		counts = append(counts, count)
	}
	return intmath.Max(counts...)/2 - intmath.Min(counts...)/2
}

func part1(in []string) {
	initial, instructions := processInput(in)
	pairs := stringToPairs(initial)
	for i := 0; i < 10; i++ {
		pairs = step(pairs, instructions)
	}
	fmt.Println("Part 1 answer:", elementDiff(initial, pairs))
}

func part2(in []string) {
	initial, instructions := processInput(in)
	pairs := stringToPairs(initial)
	for i := 0; i < 40; i++ {
		pairs = step(pairs, instructions)
	}
	fmt.Println("Part 2 answer:", elementDiff(initial, pairs))
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
