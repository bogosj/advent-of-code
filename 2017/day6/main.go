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

func findLoop() (idx, loopLen int) {
	in := input()
	seen := map[string]int{}
	for i := 0; ; i++ {
		k := hashKey(in)
		if _, ok := seen[k]; ok {
			idx = i
			loopLen = i - seen[k]
			return
		}
		seen[k] = i
		redistribute(in)
	}
}

func part1() {
	idx, _ := findLoop()
	fmt.Printf("It takes %d cycles to see a repeat\n", idx)
}

func part2() {
	_, l := findLoop()
	fmt.Printf("The loop length is %d\n", l)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
