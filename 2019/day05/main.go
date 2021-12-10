package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/computer"
)

func diagnose() int {
	c := computer.New("input.txt")
	in := make(chan int, 1)
	in <- 1
	out := c.Compute(in)
	var ret int
	for ret = range out {
	}
	return ret
}

func part1() {
	out := diagnose()
	fmt.Println("Part 1:", out)
}

func part2() {
	c := computer.New("input.txt")
	in := make(chan int, 1)
	in <- 5
	out := c.Compute(in)
	ans := <-out
	fmt.Println("Part 2:", ans)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
