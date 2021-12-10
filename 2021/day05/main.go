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

func mapFloor(in []lineSegment, considerDiagonal bool) map[intmath.Point]int {
	floor := map[intmath.Point]int{}
	for _, ls := range in {
		xStart, xEnd, yStart, yEnd := ls.start.X, ls.end.X, ls.start.Y, ls.end.Y
		xDir, yDir := 1, 1
		if xStart > xEnd {
			xDir = -1
		}
		if yStart > yEnd {
			yDir = -1
		}
		if xStart == xEnd {
			xDir = 0
		}
		if yStart == yEnd {
			yDir = 0
		}
		firstPoint := true
		if xStart == xEnd || yStart == yEnd || considerDiagonal {
			x, y := xStart, yStart
			for {
				cp := intmath.Point{X: x, Y: y}
				floor[cp]++
				x += xDir
				y += yDir
				if !firstPoint && (cp == ls.start || cp == ls.end) {
					break
				}
				firstPoint = false
			}
		}
	}
	return floor
}

func part1(in []lineSegment) {
	floor := mapFloor(in, false)
	count := 0
	for _, v := range floor {
		if v > 1 {
			count++
		}
	}
	fmt.Println("Part 1 answer:", count)
}

func part2(in []lineSegment) {
	floor := mapFloor(in, true)
	count := 0
	for _, v := range floor {
		if v > 1 {
			count++
		}
	}
	fmt.Println("Part 1 answer:", count)
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
