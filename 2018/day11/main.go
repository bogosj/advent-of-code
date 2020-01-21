package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/intmath"
)

type point = intmath.Point

const (
	input = 8561
)

func cellPower(x, y, serial int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += serial
	powerLevel *= rackID
	powerLevel /= 100
	powerLevel %= 10
	powerLevel -= 5
	return powerLevel
}

func setCellPower(x, y int, m map[point]int) {
	m[point{X: x, Y: y}] = cellPower(x, y, input)
}

func makeGrid() map[point]int {
	ret := map[point]int{}
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			setCellPower(x, y, ret)
		}
	}
	return ret
}

func totalPower(x, y int, m map[point]int) (ret int) {
	center := point{X: x + 1, Y: y + 1}
	ret += m[center]
	for _, n := range center.AllNeighbors() {
		ret += m[n]
	}
	return
}

func strongestSquare(m map[point]int) (ret point) {
	maxPower := -999999
	for x := 1; x <= 298; x++ {
		for y := 1; y <= 298; y++ {
			power := totalPower(x, y, m)
			if power > maxPower {
				maxPower = power
				ret = point{X: x, Y: y}
			}
		}
	}
	return
}

func part1() {
	p := strongestSquare(makeGrid())
	fmt.Println("The top left cell of the most powerful square is:", p)
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
