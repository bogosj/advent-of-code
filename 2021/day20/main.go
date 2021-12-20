package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func processInput(in []string) ([]bool, map[intmath.Point]bool) {
	algo := []bool{}
	image := map[intmath.Point]bool{}

	for _, c := range in[0] {
		if c == '#' {
			algo = append(algo, true)
		} else {
			algo = append(algo, false)
		}
	}

	for y, line := range in[2:] {
		for x, c := range line {
			if c == '#' {
				image[intmath.Point{X: x + 100, Y: y + 100}] = true
			}
		}
	}

	return algo, image
}

func part1(in []string) {
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
	return fileinput.ReadLines("input.txt")
}
