package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() []string {
	ret := fileinput.ReadLines("input.txt")
	sort.Strings(ret)
	return ret
}

type guard struct {
	id    string
	sleep map[int]int
}

func (g *guard) mostAsleepTime() int {
	var maxK, maxV int
	for k, v := range g.sleep {
		if v > maxV {
			maxV = v
			maxK = k
		}
	}
	return maxK
}

func (g *guard) totalSleepTime() (ret int) {
	for _, v := range g.sleep {
		ret += v
	}
	return
}

func minute(s string) int {
	f := strings.Split(s, ":")
	m := strings.ReplaceAll(f[1], "]", "")
	return intmath.Atoi(m)
}

func guards() map[string]*guard {
	ret := map[string]*guard{}
	in := input()
	var g *guard
	var found bool
	for i := 0; i < len(in); i++ {
		f := strings.Fields(in[i])
		if f[2] == "Guard" {
			if g, found = ret[f[3]]; !found {
				ng := guard{id: f[3]}
				ng.sleep = map[int]int{}
				g = &ng
				ret[f[3]] = g
			}
		} else if f[2] == "falls" {
			fMin := minute(f[1])
			i++
			f = strings.Fields(in[i])
			wMin := minute(f[1])
			for j := fMin; j < wMin; j++ {
				g.sleep[j]++
			}
		}
	}
	return ret
}

func part1() {
	var allGuards []*guard
	for _, g := range guards() {
		allGuards = append(allGuards, g)
	}
	sort.Slice(allGuards, func(i, j int) bool { return allGuards[i].totalSleepTime() < allGuards[j].totalSleepTime() })
	g := allGuards[len(allGuards)-1]
	id := intmath.Atoi(strings.ReplaceAll(g.id, "#", ""))
	answer := id * g.mostAsleepTime()
	fmt.Printf("Guard %v is asleep most at %v = %v\n", g.id, g.mostAsleepTime(), answer)
}

func part2() {
	var maxGuard *guard
	var maxK, maxV int
	for _, g := range guards() {
		for k, v := range g.sleep {
			if v > maxV {
				maxGuard = g
				maxK = k
				maxV = v
			}
		}
	}
	id := intmath.Atoi(strings.ReplaceAll(maxGuard.id, "#", ""))
	answer := id * maxK
	fmt.Printf("Guard %v is asleep most at %v = %v\n", maxGuard.id, maxK, answer)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
