package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type sectionRange struct {
	start, end int
}

func (s sectionRange) contains(other sectionRange) bool {
	return s.start <= other.start && s.end >= other.end
}

func (s sectionRange) overlaps(other sectionRange) bool {
	return (s.start <= other.end) && (s.end >= other.start)
}

func newSectionRange(s string) sectionRange {
	r := sectionRange{}
	r.start = intmath.Atoi(strings.Split(s, "-")[0])
	r.end = intmath.Atoi(strings.Split(s, "-")[1])
	return r
}

func part1(in []string) {
	count := 0
	for _, line := range in {
		r1 := newSectionRange(strings.Split(line, ",")[0])
		r2 := newSectionRange(strings.Split(line, ",")[1])
		if r1.contains(r2) || r2.contains(r1) {
			count++
		}
	}
	fmt.Printf("%d contain other ranges\n", count)
}

func part2(in []string) {
	count := 0
	for _, line := range in {
		r1 := newSectionRange(strings.Split(line, ",")[0])
		r2 := newSectionRange(strings.Split(line, ",")[1])
		if r1.overlaps(r2) {
			count++
		}
	}
	fmt.Printf("%d overlap other ranges\n", count)
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
