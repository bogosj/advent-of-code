package main

import (
	"fmt"
	"time"
)

var (
	input = []int{1, 1, 1, 3, 1, 2, 2, 1, 1, 3}
)

func lookAndSay(in []int) (ret []int) {
	last := in[0]
	count := 1
	for i := 1; i < len(in); i++ {
		next := in[i]
		if next != last {
			ret = append(ret, count, last)
			last = next
			count = 1
			continue
		}
		count++
	}
	ret = append(ret, count, last)
	return
}

func part1() {
	start := input
	for i := 0; i < 40; i++ {
		start = lookAndSay(start)
	}
	fmt.Println("Length:", len(start))
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
