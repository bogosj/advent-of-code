package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/intmath"
)

const (
	row    = 2978
	column = 3083
)

var (
	cache = map[intmath.Point]int{}
)

func valAt(r, c int) int {
	if v, ok := cache[intmath.Point{Y: r, X: c}]; ok {
		return v
	}
	if r == 1 && c == 1 {
		return 20151125
	}
	prevR := r + 1
	prevC := c - 1
	if prevC == 0 {
		prevC = prevR - 2
		prevR = 1
	}
	prevVal := valAt(prevR, prevC)
	return (prevVal * 252533) % 33554393
}

func part1() {
	for r := 100; r <= row; r++ {
		cache[intmath.Point{Y: r, X: column}] = valAt(r, column)
	}
	fmt.Println("Value is", valAt(row, column))
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
