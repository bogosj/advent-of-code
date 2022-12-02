package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

// Rock 1 A/X Paper 2 B/Y Scissors 3 C/Z
// Outcome 0 loss, 3 tie, 6 win

func part1(in []string) {
	scores := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}
	total := 0
	for _, line := range in {
		total += scores[line]
	}
	fmt.Printf("Total score: %d\n", total)
}

// Rock 1 A Paper 2 B Scissors 3 C
// X lose, Y draw, Z win
// Outcome 0 loss, 3 tie, 6 win

func part2(in []string) {
	scores := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}
	total := 0
	for _, line := range in {
		total += scores[line]
	}
	fmt.Printf("Total score: %d\n", total)
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

func input() []string {
	return fileinput.ReadLines("input.txt")
}
