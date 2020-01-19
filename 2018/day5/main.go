package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() []rune {
	lines := fileinput.ReadLines("input.txt")
	return []rune(lines[0])
}

func react(r []rune) (ret []rune) {
	for i := 0; i < len(r); i++ {
		if i == len(r)-1 {
			ret = append(ret, r[i])
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

func part1() {
	in := input()
	out := react(in)
	for len(in) != len(out) {
		in = out
		out = react(in)
	}
	fmt.Println("The length of a fully processed polymer is:", len(out))
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
