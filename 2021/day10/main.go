package main

import (
	"fmt"
	"sort"
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
	_, ok := closes[char]
	return !ok
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

func complete(line string, opens map[rune]rune) (ret string) {
	chars := []rune{}
	for _, c := range line {
		if len(chars) > 0 {
			if isOpen(c) {
				chars = append(chars, c)
			} else {
				chars = chars[:len(chars)-1]
			}
		} else {
			chars = append(chars, c)
		}
	}
	for i := len(chars) - 1; i >= 0; i-- {
		ret += string(opens[chars[i]])
	}
	return
}

func scoreIncomplete(s string) (score int) {
	for _, c := range s {
		score *= 5
		switch c {
		case ')':
			score += 1
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}
	return
}

func part2(in []string) {
	incomplete := []string{}
	opens := map[rune]rune{}
	for k, v := range closes {
		opens[v] = k
	}
	for _, line := range in {
		if badChar(line) == 'x' {
			incomplete = append(incomplete, line)
		}
	}
	scores := []int{}
	for _, line := range incomplete {
		scores = append(scores, scoreIncomplete(complete(line, opens)))
	}
	sort.Ints(scores)
	fmt.Println("Part 2 answer:", scores[len(scores)/2])
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
