package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"

	"golang.org/x/tools/container/intsets"
)

type problem struct {
	p string
}

func (p *problem) solve() int {
	eq := p.p
	ret, err := strconv.Atoi(eq)
	for err != nil {
		eq = reduce(eq)
		ret, err = strconv.Atoi(eq)
	}
	return ret
}

func solveSub(eq string) (ret int) {
	op := "+"
	for _, f := range strings.Split(eq, " ") {
		switch f {
		case "+", "*":
			op = f
		default:
			switch op {
			case "+":
				ret += intmath.Atoi(f)
			case "*":
				ret *= intmath.Atoi(f)
			}
		}

	}
	return
}

func reduce(eq string) string {
	d := maxDepth(eq)
	if d == 0 {
		return fmt.Sprint(solveSub(eq))
	}
	depth := 0
	for idx, ch := range eq {
		switch ch {
		case '(':
			depth++
		case ')':
			depth--
		}
		if depth == d {
			rest := strings.Split(eq[idx+1:], ")")
			return eq[:idx] + fmt.Sprint(solveSub(rest[0])) + strings.Join(rest[1:], ")")
		}
	}
	return ""
}

func maxDepth(eq string) int {
	c := intsets.Sparse{}
	depth := 0
	for _, ch := range eq {
		switch ch {
		case '(':
			depth++
		case ')':
			depth--
		}
		c.Insert(depth)
	}
	return c.Max()
}

func part1(in []problem) {
	total := 0
	for _, p := range in {
		total += p.solve()
	}
	fmt.Printf("All equations add up to %v\n", total)
}

func part2(in []problem) {
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

func input() []problem {
	ret := []problem{}

	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, problem{p: line})
	}

	return ret
}
