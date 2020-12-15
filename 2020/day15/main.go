package main

import (
	"fmt"
	"time"
)

func solve(count int) {
	turns := map[int][]int{
		6:  []int{5},
		0:  []int{4},
		11: []int{3},
		10: []int{2},
		1:  []int{1},
		2:  []int{0},
	}
	last := 6
	for i := 6; i < count; i++ {
		next := 0
		if len(turns[last]) > 1 {
			next = turns[last][0] - turns[last][1]
		}
		if len(turns[next]) > 0 {
			turns[next] = []int{i, turns[next][0]}
		} else {
			turns[next] = []int{i}
		}
		last = next
	}
	fmt.Println(last)
}

func main() {
	start := time.Now()
	solve(2020)
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	solve(30000000)
	fmt.Println("Part 2 done in", time.Since(start))
}
