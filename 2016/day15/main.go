package main

import (
	"fmt"
	"time"
)

type disc struct {
	order, positions, startPosition int
}

func discs() []disc {
	return []disc{
		{1, 13, 11},
		{2, 5, 0},
		{3, 17, 11},
		{4, 3, 0},
		{5, 7, 2},
		{6, 19, 17},
	}
}

func firstTime(ds []disc) int {
OUTER:
	for t := 0; ; t++ {
		for _, d := range ds {
			if (t+d.order+d.startPosition)%d.positions != 0 {
				continue OUTER
			}
		}
		return t
	}
}

func part1() {
	fmt.Println("The first time you can drop and succeed is:", firstTime(discs()))
}

func part2() {
	ds := append(discs(), disc{7, 11, 0})
	fmt.Println("The first time you can drop (with the new disc) and succeed is:", firstTime(ds))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
