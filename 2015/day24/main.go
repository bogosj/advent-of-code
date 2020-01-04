package main

import (
	"fmt"
	"math"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"

	"gonum.org/v1/gonum/stat/combin"
)

func input() (ret []int) {
	lines := fileinput.ReadLines("input.txt")
	for _, line := range lines {
		ret = append(ret, intmath.Atoi(line))
	}
	return ret
}

func calcMinQEForGroup(set []int, groups int) int {
	minQE := math.MaxInt64
	target := intmath.Sum(set...) / groups
	testLen := 1
	sumFound := false

	for !sumFound {
		combos := combin.Combinations(len(set), testLen)
		for _, combo := range combos {
			var vals []int
			for _, idx := range combo {
				vals = append(vals, set[idx])
			}
			if intmath.Sum(vals...) == target {
				qe := intmath.Product(vals...)
				if qe < minQE {
					minQE = qe
				}
				sumFound = true
			}
		}
		testLen++
	}
	return minQE
}

func calcMinQE(groupCount int) int {
	pkgs := input()
	return calcMinQEForGroup(pkgs, groupCount)
}

func part1() {
	fmt.Println("Minumum QE with three groups is", calcMinQE(3))
}

func part2() {
	fmt.Println("Minumum QE with four groups is", calcMinQE(4))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
