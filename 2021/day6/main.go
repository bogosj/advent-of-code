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

func runGenerations(fish map[int]int, generations int) map[int]int {
	for i := 0; i < generations; i++ {
		newCounter := map[int]int{}
		for i := 1; i <= 8; i++ {
			newCounter[i-1] = fish[i]
		}
		newCounter[8] = fish[0]
		newCounter[6] += fish[0]
		fish = newCounter
	}
	return fish
}

func getPopulation(fish map[int]int) (sum int) {
	for _, v := range fish {
		sum += v
	}
	return
}

func part1(in []int) {
	fish := makeCounter(in)
	fish = runGenerations(fish, 80)
	fmt.Println("Part 1 answer:", getPopulation(fish))
}

func part2(in []int) {
	fish := makeCounter(in)
	fish = runGenerations(fish, 256)
	fmt.Println("Part 2 answer:", getPopulation(fish))
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
