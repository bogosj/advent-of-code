package main

import (
	"fmt"
	"github.com/bogosj/advent-of-code/intmath"
	"time"
)

const (
	goal    = 33100000
	tooHigh = 3310000
)

func presentsAtHouse(house int) (ret int) {
	for _, i := range intmath.Factors(house) {
		ret += i * 10
	}
	return
}

func presentsAtHouse2(house int) (ret int) {
	for _, i := range intmath.Factors(house) {
		if i*50 < house {
			continue
		}
		ret += i * 11
	}
	return
}

func part1() {
	for i := 1; i < tooHigh; i++ {
		p := presentsAtHouse(i)
		if p >= goal {
			fmt.Println("House", i, "meets the goal.")
			return
		}
	}
}

func part2() {
	for i := 1; i < tooHigh; i++ {
		p := presentsAtHouse2(i)
		if p >= goal {
			fmt.Println("House", i, "meets the goal.")
			return
		}
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
