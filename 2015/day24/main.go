package main

import (
	"fmt"
	"math"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() (ret []int) {
	lines := fileinput.ReadLines("input.txt")
	for _, line := range lines {
		ret = append(ret, intmath.Atoi(line))
	}
	return ret
}

func combinations(vals []int, groups int) <-chan []int {
	o := make(chan []int)
	go func(set []int, out chan<- []int) {
		defer close(out)
		target := intmath.Sum(set...) / groups
		length := uint(len(set))
		minLen := len(set)

		// Go through all possible combinations of objects
		// from 1 (only first object in subset) to 2^length (all objects in subset)
		for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
			var subset []int

			for object := uint(0); object < length; object++ {
				// checks if object is contained in subset
				// by checking if bit 'object' is set in subsetBits
				if (subsetBits>>object)&1 == 1 {
					// add object to subset
					subset = append(subset, set[object])
				}
			}
			if len(subset) <= minLen && intmath.Sum(subset...) == target {
				out <- subset
				if minLen > len(subset) {
					minLen = len(subset)
				}
			}
		}

	}(vals, o)
	return o
}

func calcMinQE(groupCount int) int {
	pkgs := input()
	valid := combinations(pkgs, groupCount)
	var groups [][]int
	for p := range valid {
		groups = append(groups, p)
	}
	minLen := len(pkgs)
	for _, g := range groups {
		if len(g) < minLen {
			minLen = len(g)
		}
	}
	minQE := math.MaxInt64
	for _, g := range groups {
		if len(g) == minLen && intmath.Product(g...) < minQE {
			minQE = intmath.Product(g...)
		}
	}
	return minQE
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
