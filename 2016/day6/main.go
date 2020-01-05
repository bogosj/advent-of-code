package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func input() []string {
	lines := fileinput.ReadLines("input.txt")
	return lines
}

func part1() {
	counter := map[int]map[rune]int{}
	in := input()
	for i := 0; i < len(in[0]); i++ {
		counter[i] = map[rune]int{}
		for _, line := range in {
			counter[i][rune(line[i])]++
		}
	}
	fmt.Println(counter)
	fmt.Print("The message is: ")
	for i := 0; i < len(in[0]); i++ {
		var maxChar rune
		maxCount := 0
		for k, v := range counter[i] {
			if v > maxCount {
				maxCount = v
				maxChar = k
			}
		}
		fmt.Print(string(maxChar))
	}
	fmt.Println()
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
