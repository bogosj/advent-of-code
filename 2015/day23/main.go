package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2015/day23/computer"
)

func part1() {
	c := computer.New("input.txt")
	c.Run()
	fmt.Println("The value of register B is", c.Reg["b"])
}

func part2() {
	c := computer.New("input.txt")
	c.Reg["a"] = 1
	c.Run()
	fmt.Println("The value of register B is", c.Reg["b"], "if register A starts at 1.")
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
