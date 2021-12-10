package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

var closes = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
	'>': '<',
}

func isOpen(char rune) bool {
	switch char {
	case '(', '{', '[', '<':
		return true
	default:
		return false
	}
}

func badChar(line string) rune {
	chars := []rune{}
	for _, c := range line {
		if len(chars) > 0 {
			if isOpen(c) {
				chars = append(chars, c)
			} else {
				if chars[len(chars)-1] != closes[c] {
					return c
				} else {
					chars = chars[:len(chars)-1]
				}
			}
		} else {
			chars = append(chars, c)
		}
	}
	return 'x'
}

func part1(in []string) {
	b := map[rune]int{}
	for _, line := range in {
		b[badChar(line)]++
	}
	score := b[')']*3 + b[']']*57 + b['}']*1197 + b['>']*25137
	fmt.Println("Part 1 answer:", score)
}

func part2(in []string) {
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
