package main

import (
	"fmt"
	"time"
)

// Returns the next cup
func playRound(cups []int, curr, max int) (ret int) {
	// Extract next three cups
	a := cups[curr]
	b := cups[a]
	c := cups[b]
	cups[curr] = cups[c]

	// Return the next cup later
	ret = cups[c]

	// Get destination cup
	dest := curr - 1
	if dest == 0 {
		dest = max
	}
	for dest == a || dest == b || dest == c {
		dest--
		if dest == 0 {
			dest = max
		}
	}

	// Insert a/b/c after destination
	end := cups[dest]
	cups[dest] = a
	cups[c] = end
	return
}

func makeCups(in []int) []int {
	ret := make([]int, len(in)+1)
	for idx, i := range in {
		ret[i] = in[(idx+1)%len(in)]
	}
	return ret
}

func part1(in []int) {
	m := makeCups(in)
	curr := 3
	for i := 0; i < 100; i++ {
		curr = playRound(m, curr, 9)
	}
	next := m[1]
	for next != 1 {
		fmt.Print(next)
		next = m[next]
	}
	fmt.Println()
}

func expand(in []int) (ret []int) {
	ret = append(ret, in...)
	for i := 10; i <= 1_000_000; i++ {
		ret = append(ret, i)
	}
	return
}

func part2(in []int) {
	in = expand(in)
	m := makeCups(in)
	curr := 3
	for i := 0; i < 10_000_000; i++ {
		curr = playRound(m, curr, 1_000_000)
	}
	a := m[1]
	b := m[a]
	fmt.Println(a * b)
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
	return []int{3, 1, 5, 6, 7, 9, 8, 2, 4}
}
