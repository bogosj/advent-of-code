package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type lineSegment struct {
	start, end intmath.Point
}

func part1(in []lineSegment) {
	fmt.Println(in)
}

func part2(in []lineSegment) {
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

func input() []lineSegment {
	ret := []lineSegment{}

	for _, line := range fileinput.ReadLines("input.txt") {
		f := strings.Fields(line)
		ls := lineSegment{}
		ls.start = intmath.Point{
			X: intmath.Atoi(strings.Split(f[0], ",")[0]),
			Y: intmath.Atoi(strings.Split(f[0], ",")[1]),
		}
		ls.end = intmath.Point{
			X: intmath.Atoi(strings.Split(f[2], ",")[0]),
			Y: intmath.Atoi(strings.Split(f[2], ",")[1]),
		}
		ret = append(ret, ls)
	}

	return ret
}
