package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() (ret []int) {
	lines := fileinput.ReadLines("input.txt")
	for _, i := range strings.Fields(lines[0]) {
		ret = append(ret, intmath.Atoi(i))
	}
	return
}

func hashKey(in []int) string {
	var s []string
	for _, i := range in {
		s = append(s, fmt.Sprintf("%d", i))
	}
	return strings.Join(s, "|")
}

func redistribute(in []int) {
	var maxI, maxV int
	for i, v := range in {
		if v > maxV {
			maxI = i
			maxV = v
		}
	}
	blocks := in[maxI]
	in[maxI] = 0
	for i, l := maxI+1, len(in); blocks > 0; i++ {
		in[i%l]++
		blocks--
	}
}

func findLoop() (ret int) {
	in := input()
	seen := map[string]bool{}
	for i := 0; ; i++ {
		k := hashKey(in)
		if seen[k] {
			return i
		}
		seen[k] = true
		redistribute(in)
	}
}

func part1() {
	fmt.Printf("It takes %d cycles to see a repeat\n", findLoop())
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
