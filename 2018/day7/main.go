package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

type step struct {
	name    string
	preReqs []*step
}

func (s *step) String() string {
	ret := fmt.Sprintf("%v needs ", s.name)
	for _, p := range s.preReqs {
		ret += p.name + " "
	}
	return ret
}

func (s *step) ready(c map[string]bool) bool {
	for _, p := range s.preReqs {
		if !c[p.name] {
			return false
		}
	}
	return true
}

func input() map[string]*step {
	s := map[string]*step{}
	for _, line := range fileinput.ReadLines("input.txt") {
		var pre, post *step
		f := strings.Fields(line)
		if pre = s[f[1]]; pre == nil {
			pre = &step{name: f[1]}
			s[f[1]] = pre
		}
		if post = s[f[7]]; post == nil {
			post = &step{name: f[7]}
			s[f[7]] = post
		}
		post.preReqs = append(post.preReqs, pre)
	}
	return s
}

func executionOrder() (order string) {
	completed := map[string]bool{}
	allSteps := input()
	var stepNames []string
	for k := range allSteps {
		stepNames = append(stepNames, k)
	}
	sort.Strings(stepNames)
	for {
		newStepFound := false
		for _, name := range stepNames {
			if !completed[name] && allSteps[name].ready(completed) {
				completed[name] = true
				order += name
				newStepFound = true
				break
			}
		}
		if !newStepFound {
			return
		}
	}
}

func part1() {
	fmt.Println("The execution order is:", executionOrder())
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
