package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type bag struct {
	name  string
	count int
}

type rule struct {
	name string
	bags []bag
}

func parentBags(in []rule) map[string][]string {
	ret := map[string][]string{}
	for _, r := range in {
		for _, b := range r.bags {
			ret[b.name] = append(ret[b.name], r.name)
		}
	}
	return ret
}

func ruleMap(in []rule) map[string]rule {
	ret := map[string]rule{}
	for _, r := range in {
		ret[r.name] = r
	}
	return ret
}

func part1(in []rule) {
	p := parentBags(in)
	s := []string{"shiny gold"}
	pos := map[string]bool{}
	for len(s) > 0 {
		cur := s[0]
		s = s[1:]
		for _, b := range p[cur] {
			pos[b] = true
			s = append(s, b)
		}
	}
	fmt.Printf("There are %v possible outer bags.\n", len(pos))
}

func part2(in []rule) {
	m := ruleMap(in)
	c := 0
	children := m["shiny gold"].bags
	for len(children) > 0 {
		cur := children[0]
		children = children[1:]
		c += cur.count
		for i := 0; i < cur.count; i++ {
			children = append(children, m[cur.name].bags...)
		}
	}
	fmt.Printf("shiny golden bags require %v sub bags.\n", c)
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

func ruleFromLine(l string) rule {
	ret := rule{}
	f := strings.Split(l, " ")
	ret.name = strings.Join(f[0:2], " ")
	f = f[4:]
	for len(f) >= 4 {
		b := bag{
			count: intmath.Atoi(f[0]),
			name:  strings.Join(f[1:3], " "),
		}
		ret.bags = append(ret.bags, b)
		f = f[4:]
	}
	return ret
}

func input() []rule {
	ret := []rule{}

	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, ruleFromLine(line))
	}

	return ret
}
