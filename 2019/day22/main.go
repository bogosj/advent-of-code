package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/day22/deck"
)

func part1() {
	d := deck.New(10007)
	d.RunInstructions("input.txt")
	fmt.Printf("Card at position %d is %d\n", 2019, d.PosOfCard(2019))
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}
