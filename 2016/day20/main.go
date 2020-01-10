package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
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

func mergedRanges(in []ipRange) []ipRange {
	var mergedRanges []ipRange
	sort.Slice(in, func(i, j int) bool { return in[i].start < in[j].start })
	for len(in) > 0 {
		if len(in) == 1 {
			mergedRanges = append(mergedRanges, in[0])
			in = nil
			continue
		}
		if in[0].overlaps(in[1]) {
			nr := in[0].mergeWith(in[1])
			in = append([]ipRange{nr}, in[2:]...)
		} else {
			mergedRanges = append(mergedRanges, in[0])
			in = in[1:]
		}
	}
	return mergedRanges
}

func part1() {
	mr := mergedRanges(input())
	fmt.Println("The lowest valued IP is:", mr[0].end+1)
}

func part2() {
	mr := mergedRanges(input())
	ips := 0
	for i := 0; i < len(mr)-1; i++ {
		ips += mr[i+1].start - mr[i].end - 1
	}
	r := mr[len(mr)-1]
	if r.end < 4294967295 {
		ips += 4294967295 - r.end - 1
	}
	fmt.Println(ips, "IPs are allowed by the blocklist.")
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
