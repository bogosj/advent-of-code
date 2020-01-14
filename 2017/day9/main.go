package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func input() []rune {
	lines := fileinput.ReadLines("input.txt")
	return []rune(lines[0])
}

func processInput(stream []rune) (ret int) {
	var depth int
	var inGarbage bool
	for i := 0; i < len(stream); i++ {
		if inGarbage {
			switch stream[i] {
			case '>':
				inGarbage = false
			case '!':
				i++
			}
			continue
		}
		switch stream[i] {
		case '{':
			depth++
			ret += depth
		case '}':
			depth--
		case '<':
			inGarbage = true
		case '!':
			i++
		}
	}
	return
}

func part1() {
	fmt.Println("The score for all groups in the input is:", processInput(input()))
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
