package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

type group struct {
	answers []string
}

func (g *group) count() int {
	m := map[rune]int{}
	for _, a := range g.answers {
		for _, c := range a {
			m[c]++
		}
	}
	return len(m)
}

func part1(in []group) {
	c := 0
	for _, g := range in {
		c += g.count()
	}
	fmt.Printf("The sum of answers is %v\n", c)
}

func part2(in []group) {
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

func input() []group {
	ret := []group{}

	g := group{}
	for _, line := range fileinput.ReadLines("input.txt") {
		if strings.TrimSpace(line) == "" {
			ret = append(ret, g)
			g = group{}
			continue
		}
		g.answers = append(g.answers, line)
	}
	ret = append(ret, g)

	return ret
}
