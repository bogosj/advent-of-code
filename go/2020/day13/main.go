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

type bus struct {
	period, offset int
}

func (b *bus) inPosition(t int) bool {
	return intmath.Mod(t+b.offset, b.period) == 0
}

func offsets(in []string) (ret []bus) {
	for i, b := range in {
		if b != "x" {
			ret = append(ret, bus{period: intmath.Atoi(b), offset: i})
		}
	}
	return
}

func departureTime(in []string) int {
	o := offsets(in)
	timeWarp := 1
	t := 0
	toFind := map[bus]bool{}
	for _, b := range o {
		toFind[b] = true
	}
	for {
		for b := range toFind {
			if b.inPosition(t) {
				timeWarp *= b.period
				delete(toFind, b)
			}
		}
		if len(toFind) == 0 {
			break
		}
		t += timeWarp
	}
	return t
}

func part2(in schedule) {
	t := departureTime(in.buses)
	fmt.Printf("The t where all buses are offset correctly is %v\n", t)
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
