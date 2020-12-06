package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/intmath"

	"github.com/bogosj/advent-of-code/fileinput"
)

type boardingPass struct {
	data string
}

func (b *boardingPass) row() int {
	min := 0
	max := 127
	mv := 64
	s := strings.ReplaceAll(b.data, "R", "")
	s = strings.ReplaceAll(s, "L", "")
	for _, c := range s {
		if c == 'F' {
			max -= mv
		} else {
			min += mv
		}
		mv /= 2
	}
	return min
}

func (b *boardingPass) col() int {
	min := 0
	max := 7
	mv := 4
	s := strings.ReplaceAll(b.data, "F", "")
	s = strings.ReplaceAll(s, "B", "")
	for _, c := range s {
		if c == 'L' {
			max -= mv
		} else {
			min += mv
		}
		mv /= 2
	}
	return min
}

func (b *boardingPass) id() int {
	return b.row()*8 + b.col()
}

func part1(in []boardingPass) {
	max := 0
	for _, b := range in {
		max = intmath.Max(max, b.id())
	}
	fmt.Printf("The max ID is: %v\n", max)
}

func part2(in []boardingPass) {
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []boardingPass {
	ret := []boardingPass{}

	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, boardingPass{data: line})
	}

	return ret
}
