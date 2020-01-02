package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/intmath"
	combinations "github.com/mxschmitt/golang-combinations"
)

var (
	jugs = []string{"33", "14", "18", "20", "45", "35", "16", "35", "1", "13", "18", "13", "50", "44", "48", "6", "24", "41", "30", "42"}
)

func workableCombinations() (ret [][]string) {
	for _, c := range combinations.All(jugs) {
		i := 0
		for _, jug := range c {
			i += intmath.Atoi(jug)
		}
		if i == 150 {
			ret = append(ret, c)
		}
	}
	return
}

func part1() {
	fmt.Println("Combinations:", len(workableCombinations()))
}

func part2() {
	all := workableCombinations()
	minLen := len(jugs)
	for _, c := range all {
		if len(c) < minLen {
			minLen = len(c)
		}
	}
	count := 0
	for _, c := range all {
		if len(c) == minLen {
			count++
		}
	}
	fmt.Println("Combinations of min len:", count)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
