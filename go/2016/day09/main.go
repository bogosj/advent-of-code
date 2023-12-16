package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() string {
	lines := fileinput.ReadLines("input.txt")
	return lines[0]
}

func cmpSize(s string) (int, int) {
	f := strings.FieldsFunc(s, func(r rune) bool {
		return r == 'x'
	})
	return intmath.Atoi(f[0]), intmath.Atoi(f[1])
}

func decompressLen(msg string, recursive bool) int {
	if len(msg) == 0 {
		return 0
	}
	if msg[0] == '(' {
		idx := strings.Index(msg, ")")
		comp := msg[1:idx]
		amt, times := cmpSize(comp)
		if recursive {
			return times*decompressLen(msg[idx+1:idx+1+amt], recursive) + decompressLen(msg[idx+1+amt:], recursive)
		}
		return amt*times + decompressLen(msg[idx+1+amt:], recursive)
	}
	return 1 + decompressLen(msg[1:], recursive)
}

func part1() {
	d := decompressLen(input(), false)
	fmt.Println("Decompressed length:", d)
}

func part2() {
	d := decompressLen(input(), true)
	fmt.Println("Decompressed (recursive) length:", d)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
