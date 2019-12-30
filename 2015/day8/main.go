package main

import (
	"fmt"
	"github.com/bogosj/advent-of-code/fileinput"
	"time"
)

func stringWeight(s string) (weight, memory int) {
	for i := 1; i < len(s)-1; i++ {
		r := s[i]
		memory++
		switch r {
		case '\\':
			if s[i+1] == 'x' {
				weight += 4
				i += 3
			} else {
				weight += 2
				i++
			}
		default:
			weight++
		}
	}
	weight += 2
	return
}

func encodeLines(lines ...string) (ret []string) {
	for _, line := range lines {
		s := "\""
		for _, r := range line {
			switch r {
			case '"':
				s += "\\\""
			case '\\':
				s += "\\\\"
			default:
				s += string(r)
			}
		}
		s += "\""
		ret = append(ret, s)
	}
	return
}

func part1() {
	lines := fileinput.ReadLines("input.txt")
	var weight, memory int
	for _, line := range lines {
		w, m := stringWeight(line)
		weight += w
		memory += m
	}
	fmt.Printf("Weight (%d) - Memory (%d) = %d\n", weight, memory, weight-memory)
}

func part2() {
	lines := fileinput.ReadLines("input.txt")
	encodedLines := encodeLines(lines...)
	var weight, memory int
	for _, line := range encodedLines {
		w, m := stringWeight(line)
		weight += w
		memory += m
	}
	fmt.Printf("Weight (%d) - Memory (%d) = %d\n", weight, memory, weight-memory)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
