package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type monad struct {
	instructions []string
}

func newMonad(in []string) monad {
	ret := monad{}
	ret.instructions = in
	return ret
}

type rule struct {
	id, val int
}

func (m monad) rules() []rule {
	ret := []rule{}
	ruleId := 1
	r := rule{}
	constId := -1
	for i, line := range m.instructions {
		switch i % 18 {
		case 4:
			r.id = ruleId
			ruleId++
			switch strings.Fields(line)[2] {
			case "1":
				constId = 15
			case "26":
				constId = 5
			}
		case constId:
			r.val = intmath.Atoi(strings.Fields(line)[2])
			ret = append(ret, r)
			r = rule{}
		}
	}
	return ret
}

func (m monad) String() (ret string) {
	return fmt.Sprint(m.rules())
}

func part1(in []string) {
	m := newMonad(in)
	rules := []rule{}
	number := make([]string, 15)
	for _, r := range m.rules() {
		if r.val > 0 {
			rules = append(rules, r)
		} else {
			pop := rules[len(rules)-1]
			rules = rules[:len(rules)-1]
			c := r.val + pop.val
			if c > 0 {
				number[pop.id] = fmt.Sprint(9 - c)
				number[r.id] = "9"
			} else {
				number[pop.id] = "9"
				number[r.id] = fmt.Sprint(9 + c)
			}
		}
	}
	fmt.Println("Part 1 answer:", strings.Join(number, ""))
}

func part2(in []string) {
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

func input() []string {
	return fileinput.ReadLines("input.txt")
}
