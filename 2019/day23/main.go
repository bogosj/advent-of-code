package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/day23/network"
)

func part1() {
	n := network.New(50, "input.txt")
	fmt.Println("Address 255 received", n.Run(true))
}

func part2() {
	n := network.New(50, "input.txt")
	fmt.Println("Address 255 received", n.Run(false))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}
