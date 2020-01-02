package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2015/day19/medicine"
)

func part1() {
	m := medicine.New("input.txt")
	c := m.Calibrate()
	fmt.Println("Machine calibration:", c)
}

/*
This doesn't actually halt in a reasonable amount of time. It sometimes gets to the
right answer super quick, probably depending on which "Ar" key gets picked first.
*/
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
