package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2015/day7/wires"
)

func part1() {
	w := wires.New()
	w.Load("input.txt")
	w.Apply()
	fmt.Println(w.ValueOf("a"))
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
