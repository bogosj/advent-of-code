package main

import (
	"fmt"
	"time"
	"unicode"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() []rune {
	lines := fileinput.ReadLines("input.txt")
	return []rune(lines[0])
}

func react(r []rune, t rune) (ret []rune) {
	for i := 0; i < len(r); i++ {
		if i == len(r)-1 {
			ret = append(ret, r[i])
			continue
		}
		if t != '*' && unicode.ToLower(r[i]) == t {
			continue
		}
		if intmath.Abs(int(r[i])-int(r[i+1])) == 32 {
			i++
			continue
		}
		ret = append(ret, r[i])
	}
	return
}

func fullyProcessedLen(in []rune) int {
	out := react(in, '*')
	for len(in) != len(out) {
		in = out
		out = react(in, '*')
	}
	return len(out)
}

func part1() {
	l := fullyProcessedLen(input())
	fmt.Println("The length of a fully processed polymer is:", l)
}

func part2() {
	minLen := len(input())
	for r := 'a'; r <= 'z'; r++ {
		in := react(input(), r)
		outL := fullyProcessedLen(in)
		if outL < minLen {
			minLen = outL
		}
	}
	fmt.Println("The shortest polymer you can make is:", minLen)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
