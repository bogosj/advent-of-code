package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/intmath"

	"github.com/bogosj/advent-of-code/fileinput"
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

func part2(r rules, mt *ticket, ot []ticket) {
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
