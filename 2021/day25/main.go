package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func resetState(in [][]rune, from, to rune) {
	for y := range in {
		for x := range in[y] {
			if in[y][x] == from {
				in[y][x] = to
			}
		}
	}
}

func step(in [][]rune) bool {
	somethingMoved := false
	maxY := len(in)
	maxX := len(in[0])
	// east
	for y := range in {
		for x := range in[y] {
			if in[y][x] == '>' {
				if in[y][(x+1)%maxX] == '.' {
					in[y][x] = 'X'
					in[y][(x+1)%maxX] = '<'
					somethingMoved = true
				}
			}
		}
	}
	resetState(in, 'X', '.')
	// south
	for y := range in {
		for x := range in[y] {
			if in[y][x] == 'v' {
				if in[(y+1)%maxY][x] == '.' {
					in[y][x] = 'X'
					in[(y+1)%maxY][x] = '^'
					somethingMoved = true
				}
			}
		}
	}
	resetState(in, 'X', '.')
	resetState(in, '^', 'v')
	resetState(in, '<', '>')
	return somethingMoved
}

func toField(in []string) [][]rune {
	ret := [][]rune{}
	for _, line := range in {
		ret = append(ret, []rune(line))
	}
	return ret
}

func countSteps(in []string) int {
	current := toField(in)
	count := 0
	for {
		count++
		if !step(current) {
			break
		}
	}
	return count
}

func part1(in []string) {
	fmt.Println("Part 1 answer:", countSteps(in))
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
}

func input() []string {
	return fileinput.ReadLines("input.txt")
}
