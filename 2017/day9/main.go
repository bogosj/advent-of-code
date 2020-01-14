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

func processInput(stream []rune) (ret, gRet int) {
	var depth int
	var inGarbage bool
	for i := 0; i < len(stream); i++ {
		if inGarbage {
			switch stream[i] {
			case '>':
				inGarbage = false
			case '!':
				i++
			default:
				gRet++
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
	s, _ := processInput(input())
	fmt.Println("The score for all groups in the input is:", s)
}

func part2() {
	_, g := processInput(input())
	fmt.Println("The number ofnon-cancelled garbage characters is:", g)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
