package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type star struct {
	a, b, c, d int
}

func (s star) distanceTo(os star) int {
	abs := intmath.Abs
	return abs(s.a-os.a) + abs(s.b-os.b) + abs(s.c-os.c) + abs(s.d-os.d)
}

type constellation struct {
	stars []star
}

func (c *constellation) mergeInto(oc *constellation) bool {
	for _, s1 := range c.stars {
		for _, s2 := range oc.stars {
			if s1.distanceTo(s2) <= 3 {
				oc.stars = append(oc.stars, c.stars...)
				return true
			}
		}
	}
	return false
}

func input() (ret []star) {
	atoi := intmath.Atoi
	for _, line := range fileinput.ReadLines("input.txt") {
		f := strings.Split(line, ",")
		ret = append(ret, star{
			a: atoi(f[0]),
			b: atoi(f[1]),
			c: atoi(f[2]),
			d: atoi(f[3]),
		})
	}
	return
}

func part1() {
	allStars := input()
	var universe []*constellation
	for _, s := range allStars {
		universe = append(universe, &constellation{stars: []star{s}})
	}

MERGE:
	for {
		for i := 0; i < len(universe); i++ {
			for j := i + 1; j < len(universe); j++ {
				if universe[i].mergeInto(universe[j]) {
					universe = append(universe[:i], universe[i+1:]...)
					continue MERGE
				}
			}
		}
		break
	}
	fmt.Println(len(universe))
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
