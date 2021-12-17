package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/intmath"
)

func shoot(xv, yv int, topLeft, bottomRight intmath.Point) (hit bool, max int) {
	curr := intmath.Point{}
	for curr.X < bottomRight.X && curr.Y > bottomRight.Y {
		curr.X += xv
		curr.Y += yv
		if curr.Y > max {
			max = curr.Y
		}
		if xv > 0 {
			xv--
		}
		if xv < 0 {
			xv++
		}
		yv--
		if curr.ContainedIn(topLeft, bottomRight) {
			return true, max
		}
	}
	return false, 0
}

func part1(topLeft, bottomRight intmath.Point) {
	maxY := 0
	for xv := 1; xv < bottomRight.X; xv++ {
		for yv := 1; yv < 1000; yv++ {
			hit, max := shoot(xv, yv, topLeft, bottomRight)
			if hit && max > maxY {
				maxY = max
			}
		}
	}
	fmt.Println("Part 1 answer:", maxY)
}

func part2(topLeft, bottomRight intmath.Point) {
}

func main() {
	topLeft, bottomRight := input()
	start := time.Now()
	part1(topLeft, bottomRight)
	fmt.Println("Part 1 done in", time.Since(start))
	topLeft, bottomRight = input()
	start = time.Now()
	part2(topLeft, bottomRight)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() (topLeft, bottomRight intmath.Point) {
	/*
		// sample
		topLeft.X = 20
		topLeft.Y = -5
		bottomRight.X = 30
		bottomRight.Y = -10
	*/
	// target
	topLeft.X = 287
	topLeft.Y = -48
	bottomRight.X = 309
	bottomRight.Y = -76

	return
}
