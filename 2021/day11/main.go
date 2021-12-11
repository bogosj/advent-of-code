package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func countFlashes(in map[intmath.Point]int) (flashes int) {
	// Increase level for all
	for k, v := range in {
		in[k] = v + 1
	}

	for {
		flashed := false
		for k, v := range in {
			if v > 9 {
				flashed = true
				for _, n := range k.AllNeighbors() {
					v, ok := in[n]
					if ok && v != -1 {
						in[n] = v + 1
					}
				}
				in[k] = -1
			}
		}
		if !flashed {
			break
		}
	}

	// Reset flashed
	for k, v := range in {
		if v == -1 {
			in[k] = 0
			flashes++
		}
	}
	return
}

func printGrid(in map[intmath.Point]int, xm, ym int) {
	for y := 0; y < ym; y++ {
		for x := 0; x < xm; x++ {
			fmt.Print(in[intmath.Point{X: x, Y: y}])
		}
		fmt.Println()
	}
	fmt.Println()
}

func part1(in map[intmath.Point]int) {
	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += countFlashes(in)
	}
	fmt.Println("Part 1 answer:", flashes)
}

func part2(in map[intmath.Point]int) {
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

func input() map[intmath.Point]int {
	ret := map[intmath.Point]int{}

	for y, line := range fileinput.ReadLines("input.txt") {
		for x, c := range line {
			ret[intmath.Point{X: x, Y: y}] = int(c) - 48
		}
	}

	return ret
}
