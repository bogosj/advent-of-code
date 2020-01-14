package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() []string {
	lines := fileinput.ReadLines("input.txt")
	return strings.FieldsFunc(lines[0], func(r rune) bool { return r == ',' })
}

func countSteps(in []string) int {
	s := map[string]int{}
	for _, dir := range in {
		s[dir]++
	}
	for _, pair := range [][]string{{"n", "s"}, {"ne", "sw"}, {"nw", "se"}} {
		min := intmath.Min(s[pair[0]], s[pair[1]])
		s[pair[0]] -= min
		s[pair[1]] -= min
	}
	mergePairs := [][]string{
		{"n", "ne", "nw"},
		{"s", "se", "sw"},
		{"ne", "n", "se"},
		{"nw", "n", "sw"},
		{"se", "s", "ne"},
		{"sw", "s", "nw"},
	}
	for _, merge := range mergePairs {
		min := intmath.Min(s[merge[1]], s[merge[2]])
		s[merge[0]] += min
		s[merge[1]] -= min
		s[merge[2]] -= min
	}
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}

func part1() {
	fmt.Println("Steps required:", countSteps(input()))
}

func part2() {
	var max int
	in := input()
	for i := 1522; i < len(in); i++ {
		s := countSteps(in[0:i])
		if s > max {
			max = s
		}
	}
	fmt.Println("The maximum distance was:", max)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
