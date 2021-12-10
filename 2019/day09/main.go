package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/computer"
)

func test() {
	for _, i := range []int{1, 2, 3} {
		c := computer.New(fmt.Sprintf("test%v.txt", i))
		in := make(chan int, 1)
		in <- 0
		for out := range c.Compute(in) {
			fmt.Printf("%v", out)
		}
		fmt.Println()
	}
}

func part1() {
	c := computer.New("input.txt")
	in := make(chan int, 1)
	in <- 1
	out := <-c.Compute(in)
	fmt.Println("TESTS:", out)
}

func part2() {
	c := computer.New("input.txt")
	in := make(chan int, 1)
	in <- 2
	out := <-c.Compute(in)
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
