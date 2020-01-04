package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2015/day23/computer"
)

func part1() {
	c := computer.New("input.txt")
	c.Run()
	fmt.Println("The value of register B is:", c.Reg["b"])
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
