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
				newPoint.Y = coord - (point.Y - coord)
				paper[newPoint] = true
				delete(paper, point)
			}
		}
	}
}

func part1(in []string) {
	paper := markPoints(in)
	fold(paper, "x", 655)

	count := 0
	for _, v := range paper {
		if v {
			count++
		}
	}
	fmt.Println("Part 1 answer:", count)
}

func printPaper(paper map[intmath.Point]bool) {
	xs, ys := []int{}, []int{}
	for p := range paper {
		xs = append(xs, p.X)
		ys = append(ys, p.Y)
	}
	for y := 0; y <= intmath.Max(ys...); y++ {
		for x := 0; x <= intmath.Max(xs...); x++ {
			if paper[intmath.Point{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func part2(in []string) {
	paper := markPoints(in)
	fold(paper, "x", 655)
	fold(paper, "y", 447)
	fold(paper, "x", 327)
	fold(paper, "y", 223)
	fold(paper, "x", 163)
	fold(paper, "y", 111)
	fold(paper, "x", 81)
	fold(paper, "y", 55)
	fold(paper, "x", 40)
	fold(paper, "y", 27)
	fold(paper, "y", 13)
	fold(paper, "y", 6)
	printPaper(paper)
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
