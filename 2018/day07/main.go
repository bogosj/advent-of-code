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

func executionOrder(parallel bool) (order string, seconds int) {
	completed := map[string]bool{}
	allSteps := input()
	var stepNames []string
	for k := range allSteps {
		stepNames = append(stepNames, k)
	}
	sort.Strings(stepNames)
	workersAvail := 5
	inProgress := map[string]int{}
	for {
		newStepFound := false
		for _, name := range stepNames {
			if !completed[name] && allSteps[name].ready(completed) {
				if parallel && workersAvail > 0 {
					if _, ok := inProgress[name]; ok {
						continue
					}
					workersAvail--
					inProgress[name] = 61 + int(name[0]-'A')
					fmt.Println("Queuing", name)
				} else {
					completed[name] = true
					order += name
					newStepFound = true
					break
				}
			}
		}
		if parallel {
			if workersAvail == 5 {
				return
			}
			for jc := false; !jc; {
				for k := range inProgress {
					inProgress[k]--
					if inProgress[k] == 0 {
						completed[k] = true
						delete(inProgress, k)
						jc = true
						workersAvail++
					}
				}
				seconds++
			}
		} else {
			if !newStepFound {
				return
			}
		}
	}
}

func part1() {
	o, _ := executionOrder(false)
	fmt.Println("The execution order is:", o)
}

func part2() {
	_, s := executionOrder(true)
	fmt.Printf("If executing in parallel, it takes %d seconds\n", s)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
