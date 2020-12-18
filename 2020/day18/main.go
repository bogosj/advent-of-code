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
	p        string
	advanced bool
}

func (p *problem) solve() int {
	eq := p.p
	ret, err := strconv.Atoi(eq)
	for err != nil {
		eq = reduce(eq, p.advanced)
		ret, err = strconv.Atoi(eq)
	}
	return ret
}

func solveSub(eq string, advanced bool) (ret int) {
	if advanced {
		fields := strings.Fields(eq)
		if len(fields)==1 {
			return intmath.Atoi(fields[0])
		}
		for idx, s := range fields {
			if s == "+" {
				newFields := fields[:idx-1]
				newFields = append(newFields, fmt.Sprint(intmath.Atoi(fields[idx-1])+intmath.Atoi(fields[idx+1])))
				newFields = append(newFields, fields[idx+2:]...)
				return solveSub(strings.Join(newFields, " "), advanced)
			}
		}
		for idx, s := range fields {
			if s == "*" {
				newFields := fields[:idx-1]
				newFields = append(newFields, fmt.Sprint(intmath.Atoi(fields[idx-1])*intmath.Atoi(fields[idx+1])))
				newFields = append(newFields, fields[idx+2:]...)
				return solveSub(strings.Join(newFields, " "), advanced)
			}
		}
	} else {
		op := "+"
		for _, f := range strings.Fields(eq) {
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
	}
	return
}

func reduce(eq string, advanced bool) string {
	d := maxDepth(eq)
	if d == 0 {
		return fmt.Sprint(solveSub(eq, advanced))
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
			return eq[:idx] + fmt.Sprint(solveSub(rest[0], advanced)) + strings.Join(rest[1:], ")")
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
	total := 0
	for _, p := range in {
		p.advanced = true
		total += p.solve()
	}
	fmt.Printf("All equations add up to %v\n", total)
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
