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

func mapFloor(in []lineSegment) map[intmath.Point]int {
	floor := map[intmath.Point]int{}
	for _, ls := range in {
		xStart, xEnd, yStart, yEnd := 0, 0, 0, 0
		if ls.start.X == ls.end.X {
			xStart, xEnd = ls.start.X, ls.start.X
			if ls.start.Y > ls.end.Y {
				yStart, yEnd = ls.end.Y, ls.start.Y
			} else {
				yEnd, yStart = ls.end.Y, ls.start.Y
			}
		} else if ls.start.Y == ls.end.Y {
			yStart, yEnd = ls.start.Y, ls.start.Y
			if ls.start.X > ls.end.X {
				xStart, xEnd = ls.end.X, ls.start.X
			} else {
				xEnd, xStart = ls.end.X, ls.start.X
			}
		} else {
			continue
		}
		for x := xStart; x <= xEnd; x++ {
			for y := yStart; y <= yEnd; y++ {
				floor[intmath.Point{X: x, Y: y}]++
			}
		}
	}
	return floor
}

func part1(in []lineSegment) {
	floor := mapFloor(in)
	count := 0
	for p, v := range floor {
		if v > 1 {
			count++
		}
	}
	fmt.Println("Part 1 answer:", count)
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
