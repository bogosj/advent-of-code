package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type schedule struct {
	wait  int
	buses []string
}

func (s *schedule) activeBuses() (ret []int) {
	for _, b := range s.buses {
		if b != "x" {
			ret = append(ret, intmath.Atoi(b))
		}
	}
	return
}

func part1(in schedule) {
	minT := 1000000
	minB := 0
	for _, b := range in.activeBuses() {
		arrive := (in.wait/b + 1) * b
		w := arrive - in.wait
		if w < minT {
			minT = w
			minB = b
		}
	}
	fmt.Printf("Waiting fur bus %v for %v minutes: %v\n", minB, minT, minT*minB)
}

func part2(in schedule) {
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

func input() schedule {
	ret := schedule{}
	in := fileinput.ReadLines("input.txt")
	ret.wait = intmath.Atoi(in[0])
	ret.buses = strings.Split(in[1], ",")
	return ret
}
