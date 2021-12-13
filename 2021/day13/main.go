package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func markPoints(in []string) map[intmath.Point]bool {
	ret := map[intmath.Point]bool{}
	for _, line := range in {
		if strings.HasPrefix(line, "fold") {
			continue
		}

		f := strings.Split(line, ",")
		if len(f) == 2 {
			ret[intmath.Point{X: intmath.Atoi(f[0]), Y: intmath.Atoi(f[1])}] = true
		}
	}
	return ret
}

func fold(paper map[intmath.Point]bool, dir string, coord int) {
	for point := range paper {
		if dir == "x" {
			if point.X > coord {
				newPoint := intmath.Point{Y: point.Y}
				newPoint.X = coord - (point.X - coord)
				paper[newPoint] = true
				delete(paper, point)
			}
		} else {
			if point.Y > coord {
				newPoint := intmath.Point{X: point.X}
				newPoint.Y = (2 * (point.Y - coord)) - coord
				paper[newPoint] = true
				delete(paper, point)
			}
		}
	}
}

func part1(in []string) {
	paper := markPoints(in)
	fold(paper, "x", 655)
	//fold(paper, "y", 7)

	count := 0
	for p, v := range paper {
		if v {
			count++
			fmt.Println(p)
		}
	}
	fmt.Println("Part 1 answer:", count)
}

func part2(in []string) {
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

func input() []string {
	return fileinput.ReadLines("input.txt")
}
