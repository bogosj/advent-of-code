package main

import (
	"fmt"
	"github.com/bogosj/advent-of-code/2015/day19/medicine"
	"time"
)

func part1() {
	m := medicine.New("input.txt")
	c := m.Calibrate()
	fmt.Println("Machine calibration:", c)
}

func part2() {
	m := medicine.New("input.txt")
	c := m.Build()
	fmt.Printf("Machine required %v steps to build the molecule.\n", c)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
