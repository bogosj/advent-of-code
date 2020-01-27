package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type point = intmath.Point
type state rune

const (
	clay    = '#'
	flow    = '+'
	water   = '|'
	settled = '~'
)

func input() map[point]state {
	ret := map[point]state{}
	for _, line := range fileinput.ReadLines("input.txt") {
		f := strings.Split(line, ", ")
		primary := strings.Split(f[0], "=")
		secondary := strings.Split(f[1], "=")
		sRange := strings.Split(secondary[1], "..")
		if primary[0] == "x" {
			x := intmath.Atoi(primary[1])
			for y := intmath.Atoi(sRange[0]); y <= intmath.Atoi(sRange[1]); y++ {
				ret[point{X: x, Y: y}] = clay
			}
		} else {
			y := intmath.Atoi(primary[1])
			for x := intmath.Atoi(sRange[0]); x <= intmath.Atoi(sRange[1]); x++ {
				ret[point{X: x, Y: y}] = clay
			}
		}
	}
	return ret
}

func findMinY(scan map[point]state) int {
	var ys []int
	for p := range scan {
		ys = append(ys, p.Y)
	}
	return intmath.Min(ys...)
}

func findMaxY(scan map[point]state) int {
	var ys []int
	for p := range scan {
		ys = append(ys, p.Y)
	}
	return intmath.Max(ys...)
}

func containedDir(p point, scan map[point]state, dir int) bool {
	for scan[p] != clay {
		bp := p
		scan[p] = settled
		bp.Y++
		if below, ok := scan[bp]; ok {
			if below == clay || below == settled {
				p.X += dir
				continue
			} else {
				scan[p] = flow
				return false
			}
		} else {
			scan[p] = flow
			return false
		}
	}
	return true
}

func backtrack(p point, scan map[point]state, dir int) {
	for scan[p] != clay && scan[p] != flow {
		scan[p] = water
		p.X += dir
	}
}

func contained(p point, scan map[point]state) bool {
	left := containedDir(p, scan, -1)
	right := containedDir(p, scan, 1)

	if !(left && right) {
		backtrack(p, scan, -1)
		backtrack(p, scan, 1)
	}

	return left && right
}

func nextFillPoint(scan map[point]state) (point, error) {
	for p, s := range scan {
		if s == flow {
			return p, nil
		}
	}
	return point{}, errors.New("no more fill points")
}

func updateScan(scan map[point]state) {
	maxY := findMaxY(scan)
	var p point
	var err error
FILLPOINTS:
	for {
		p, err = nextFillPoint(scan)
		if err != nil {
			return
		}
		for {
			scan[p] = water
			p.Y++
			if scan[p] == settled || scan[p] == clay || scan[p] == flow {
				break
			}
			if p.Y > maxY {
				continue FILLPOINTS
			}
		}
		p.Y--
		for contained(p, scan) {
			p.Y--
		}
	}
}

func part1() {
	scan := input()
	minY := findMinY(scan)
	scan[point{X: 500}] = flow
	updateScan(scan)
	var c int
	for p, v := range scan {
		if (v == water || v == settled) && p.Y >= minY {
			c++
		}
	}
	fmt.Println(c)
}

func part2() {
	scan := input()
	minY := findMinY(scan)
	scan[point{X: 500}] = flow
	updateScan(scan)
	var c int
	for p, v := range scan {
		if v == settled && p.Y >= minY {
			c++
		}
	}
	fmt.Println(c)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
