package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/day24/life"
)

func part1() {
	l := life.New("input.txt")
	fmt.Println("The first repeated map scores:", l.Run())
	fmt.Println("And looks like:")
	fmt.Println(l.String())
}

func part2() {
	i := life.NewInfinite("input.txt")
	i.Run(200)
	fmt.Printf("After 200 minutes there are %d bugs.\n", i.BugCount())
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}
