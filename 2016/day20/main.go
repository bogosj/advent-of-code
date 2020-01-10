package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
	"github.com/gammazero/deque"
)

type ipRange struct {
	start, end int
}

func (i ipRange) overlaps(oi ipRange) bool {
	return oi.start >= i.start && oi.start <= i.end+1
}

func (i ipRange) mergeWith(oi ipRange) ipRange {
	return ipRange{start: i.start, end: intmath.Max(i.end, oi.end)}
}

func input() (ret []ipRange) {
	lines := fileinput.ReadLines("input.txt")
	for _, line := range lines {
		f := strings.FieldsFunc(line, func(r rune) bool { return r == '-' })
		ret = append(ret, ipRange{start: intmath.Atoi(f[0]), end: intmath.Atoi(f[1])})
	}
	return
}

func mergedRanges() deque.Deque {
	var ranges, mergedRanges deque.Deque
	in := input()
	sort.Slice(in, func(i, j int) bool { return in[i].start < in[j].start })
	for _, ir := range in {
		ranges.PushFront(ir)
	}
	for ranges.Len() > 0 {
		if ranges.Len() == 1 {
			mergedRanges.PushFront(ranges.PopFront())
			continue
		}
		fmt.Println(r1, r2)
		if r1.overlaps(r2) {
			ranges.PushBack(r1.mergeWith(r2))
		} else {
			mergedRanges.PushFront(r1)
			mergedRanges.PushFront(r2)
		}
	}
	return mergedRanges
}

func part1() {
	mr := mergedRanges()
	fmt.Println("The lowest valued IP is:", mr.Back().(ipRange).end+1)
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
