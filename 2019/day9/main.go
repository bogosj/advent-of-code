package main

import (
	"fmt"
	"time"

	"jamesbogosian.com/advent-of-code/2019/computer"
)

func test() {
	for _, i := range []int{1, 2, 3} {
		c := computer.New(fmt.Sprintf("test%v.txt", i))
		out := c.Compute(0)
		fmt.Printf("%v ", out)
		fmt.Println()
	}
}

func part1() {
	c := computer.New("input.txt")
	out := c.Compute(1)
	fmt.Println("TESTS:", out)
}

func part2() {
	c := computer.New("input.txt")
	out := c.Compute(2)
	fmt.Println("Result:", out)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
