package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

var (
	ruleRegex = regexp.MustCompile(`(.*): (.*) or (.*)`)
)

type rules struct {
	r []string
}

func (r *rules) isValid(i int) bool {
	for _, r := range r.r {
		m := ruleRegex.FindStringSubmatch(r)
		for _, g := range m[2:] {
			v := strings.Split(g, "-")
			if i >= intmath.Atoi(v[0]) && i <= intmath.Atoi(v[1]) {
				return true
			}
		}
	}
	return false
}

func (r *rules) validForRule(ruleIDx, i int) bool {
	m := ruleRegex.FindStringSubmatch(r.r[ruleIDx])
	for _, g := range m[2:] {
		v := strings.Split(g, "-")
		if i >= intmath.Atoi(v[0]) && i <= intmath.Atoi(v[1]) {
			return true
		}
	}
	return false
}

type ticket struct {
	val []int
}

func part1(r rules, mt *ticket, ot []ticket) {
	invalid := []int{}
	for _, t := range ot {
		for _, i := range t.val {
			if !r.isValid(i) {
				invalid = append(invalid, i)
			}
		}
	}
	fmt.Printf("The sum of invalid fields is %v.\n", intmath.Sum(invalid...))
}

func validTickets(r rules, ot []ticket) (ret []ticket) {
OUTER:
	for _, t := range ot {
		for _, i := range t.val {
			if !r.isValid(i) {
				continue OUTER
			}
		}
		ret = append(ret, t)
	}
	return
}

func part2(r rules, mt *ticket, ot []ticket) {
	columnToRule := map[int]map[int]int{}
	for i := range r.r {
		columnToRule[i] = map[int]int{}
		for j := range r.r {
			columnToRule[i][j] = j
		}
	}

	vt := validTickets(r, ot)
	for _, t := range vt {
		for c, v := range t.val {
			for ruleID := range r.r {
				if !r.validForRule(ruleID, v) {
					delete(columnToRule[c], ruleID)
				}
			}
		}
	}

	uniquePairs := map[int]int{}
	for len(columnToRule) > 0 {
		for k, v := range columnToRule {
			if len(v) == 1 {
				for k2, v2 := range v {
					uniquePairs[k] = v2
					for _, v3 := range columnToRule {
						delete(v3, k2)
					}
				}
				delete(columnToRule, k)
			}
		}
	}

	ruleToColumn := map[int]int{}
	for k, v := range uniquePairs {
		ruleToColumn[v] = k
	}

	total := 1
	for i := 0; i < 6; i++ {
		total *= mt.val[ruleToColumn[i]]
	}
	fmt.Printf("The product of departure values is %v.\n", total)
}

func main() {
	r, mt, ot := input()
	start := time.Now()
	part1(r, mt, ot)
	fmt.Println("Part 1 done in", time.Since(start))
	r, mt, ot = input()
	start = time.Now()
	part2(r, mt, ot)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() (rules, *ticket, []ticket) {
	d := fileinput.ReadLines("input.txt")
	r := rules{}
	var mt *ticket
	ot := []ticket{}

	for _, line := range d {
		if ruleRegex.MatchString(line) {
			r.r = append(r.r, line)
		} else if strings.Contains(line, ",") {
			if mt == nil {
				mt = &ticket{}
				for _, v := range strings.Split(line, ",") {
					mt.val = append(mt.val, intmath.Atoi(v))
				}
			} else {
				t := ticket{}
				for _, v := range strings.Split(line, ",") {
					t.val = append(t.val, intmath.Atoi(v))
				}
				ot = append(ot, t)
			}
		}
	}
	return r, mt, ot
}
