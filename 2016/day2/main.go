package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

var (
	keypad = [][]rune{
		{' ', ' ', ' ', ' ', ' '},
		{' ', '1', '2', '3', ' '},
		{' ', '4', '5', '6', ' '},
		{' ', '7', '8', '9', ' '},
		{' ', ' ', ' ', ' ', ' '},
	}
	keypad2 = [][]rune{
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', '1', ' ', ' ', ' '},
		{' ', ' ', '2', '3', '4', ' ', ' '},
		{' ', '5', '6', '7', '8', '9', ' '},
		{' ', ' ', 'A', 'B', 'C', ' ', ' '},
		{' ', ' ', ' ', 'D', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},
	}
)

func input() (ret []string) {
	lines := fileinput.ReadLines("input.txt")
	return lines
}

func walk(k [][]rune, x, y int, path []string) (ret string) {
	for _, line := range path {
		for _, r := range line {
			switch r {
			case 'D':
				y++
				if k[y][x] == ' ' {
					y--
				}
			case 'U':
				y--
				if k[y][x] == ' ' {
					y++
				}
			case 'R':
				x++
				if k[y][x] == ' ' {
					x--
				}
			case 'L':
				x--
				if k[y][x] == ' ' {
					x++
				}
			}
		}
		ret += string(k[y][x])
	}
	return
}

func part1() {
	fmt.Println(walk(keypad, 2, 2, input()))
}

func part2() {
	fmt.Println(walk(keypad2, 3, 3, input()))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
