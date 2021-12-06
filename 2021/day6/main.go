package main

import (
	"fmt"
	"time"
)

func makeCounter(in []int) map[int]int {
	ret := map[int]int{}
	for _, i := range in {
		ret[i]++
	}
	return ret
}

func part1(in []int) {
	counter := makeCounter(in)
	for i := 0; i < 80; i++ {
		newCounter := map[int]int{}
		for i := 1; i <= 8; i++ {
			newCounter[i-1] = counter[i]
		}
		newCounter[8] = counter[0]
		newCounter[6] += counter[0]
		counter = newCounter
	}
	sum := 0
	for _, v := range counter {
		sum += v
	}
	fmt.Println("Part 1 answer:", sum)
}

func part2(in []int) {
	counter := makeCounter(in)
	for i := 0; i < 256; i++ {
		newCounter := map[int]int{}
		for i := 1; i <= 8; i++ {
			newCounter[i-1] = counter[i]
		}
		newCounter[8] = counter[0]
		newCounter[6] += counter[0]
		counter = newCounter
	}
	sum := 0
	for _, v := range counter {
		sum += v
	}
	fmt.Println("Part 1 answer:", sum)
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	in = input()
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []int {
	//return []int{3, 4, 3, 1, 2}
	return []int{2, 4, 1, 5, 1, 3, 1, 1, 5, 2, 2, 5, 4, 2, 1, 2, 5, 3, 2, 4, 1, 3, 5, 3, 1, 3, 1, 3, 5, 4, 1, 1, 1, 1, 5, 1, 2, 5, 5, 5, 2, 3, 4, 1, 1, 1, 2, 1, 4, 1, 3, 2, 1, 4, 3, 1, 4, 1, 5, 4, 5, 1, 4, 1, 2, 2, 3, 1, 1, 1, 2, 5, 1, 1, 1, 2, 1, 1, 2, 2, 1, 4, 3, 3, 1, 1, 1, 2, 1, 2, 5, 4, 1, 4, 3, 1, 5, 5, 1, 3, 1, 5, 1, 5, 2, 4, 5, 1, 2, 1, 1, 5, 4, 1, 1, 4, 5, 3, 1, 4, 5, 1, 3, 2, 2, 1, 1, 1, 4, 5, 2, 2, 5, 1, 4, 5, 2, 1, 1, 5, 3, 1, 1, 1, 3, 1, 2, 3, 3, 1, 4, 3, 1, 2, 3, 1, 4, 2, 1, 2, 5, 4, 2, 5, 4, 1, 1, 2, 1, 2, 4, 3, 3, 1, 1, 5, 1, 1, 1, 1, 1, 3, 1, 4, 1, 4, 1, 2, 3, 5, 1, 2, 5, 4, 5, 4, 1, 3, 1, 4, 3, 1, 2, 2, 2, 1, 5, 1, 1, 1, 3, 2, 1, 3, 5, 2, 1, 1, 4, 4, 3, 5, 3, 5, 1, 4, 3, 1, 3, 5, 1, 3, 4, 1, 2, 5, 2, 1, 5, 4, 3, 4, 1, 3, 3, 5, 1, 1, 3, 5, 3, 3, 4, 3, 5, 5, 1, 4, 1, 1, 3, 5, 5, 1, 5, 4, 4, 1, 3, 1, 1, 1, 1, 3, 2, 1, 2, 3, 1, 5, 1, 1, 1, 4, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 2, 5, 3}
}
