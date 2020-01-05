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

func getCounts() map[int]map[rune]int {
	counter := map[int]map[rune]int{}
	in := input()
	for i := 0; i < len(in[0]); i++ {
		counter[i] = map[rune]int{}
		for _, line := range in {
			counter[i][rune(line[i])]++
		}
	}
	return counter
}

func part1() {
	in := input()
	counter := getCounts()
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
	in := input()
	counter := getCounts()
	fmt.Print("The real message is: ")
	for i := 0; i < len(in[0]); i++ {
		var minChar rune
		minCount := len(in)
		for k, v := range counter[i] {
			if v < minCount {
				minCount = v
				minChar = k
			}
		}
		fmt.Print(string(minChar))
	}
	fmt.Println()
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
