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

func totalPower(cx, cy, size int, m map[point]int) (ret int) {
	for x := cx; x < cx+size; x++ {
		for y := cy; y < cy+size; y++ {
			ret += m[point{X: x, Y: y}]
		}
	}
	return
}

func strongestSquare(m map[point]int, size int) (ret point, retPower int) {
	retPower = -999999
	for x := 1; x <= 300; x++ {
		if x+size > 300 {
			continue
		}
		for y := 1; y <= 300; y++ {
			power := totalPower(x, y, size, m)
			if power > retPower {
				retPower = power
				ret = point{X: x, Y: y}
			}
		}
	}
	return
}

func part1() {
	p, _ := strongestSquare(makeGrid(), 3)
	fmt.Println("The top left cell of the most powerful square is:", p)
}

func part2() {
	g := makeGrid()
	for size := 1; size < 300; size++ {
		p, power := strongestSquare(g, size)
		fmt.Printf("The top left cell of the most powerful square is (size %d): %v, power:%v\n", size, p, power)
	}
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
