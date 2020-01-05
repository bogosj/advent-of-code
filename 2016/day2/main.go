package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

var (
	keypad = [][]rune{
		{'1', '2', '3'},
		{'4', '5', '6'},
		{'7', '8', '9'},
	}
)

func input() (ret []string) {
	lines := fileinput.ReadLines("input.txt")
	return lines
}

func walk(path []string) (ret string) {
	x, y := 1, 1
	for _, line := range path {
		for _, r := range line {
			switch r {
			case 'D':
				if y < 2 {
					y++
				}
			case 'U':
				if y > 0 {
					y--
				}
			case 'R':
				if x < 2 {
					x++
				}
			case 'L':
				if x > 0 {
					x--
				}
			}
		}
		ret += string(keypad[y][x])
	}
	return
}

func part1() {
	fmt.Println(walk(input()))
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
