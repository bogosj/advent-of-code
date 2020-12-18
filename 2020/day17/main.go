package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

type cube struct {
	on        bool
	neighbors int
}

func solve(universe map[string]cube, dimensions int) int {
	for cycle := 0; cycle < 6; cycle++ {
		for coords, c := range universe {
			if !c.on {
				continue
			}
			var x, y, z, w int
			fmt.Sscanf(coords, "%d|%d|%d|%d", &x, &y, &z, &w)
			for xd := -1; xd <= 1; xd++ {
				for yd := -1; yd <= 1; yd++ {
					for zd := -1; zd <= 1; zd++ {
						wdMin, wdMax := -1, 1
						if dimensions == 3 {
							wdMin, wdMax = 0, 0
						}
						for wd := wdMin; wd <= wdMax; wd++ {
							if xd == 0 && yd == 0 && zd == 0 && wd == 0 {
								continue
							}
							ncoord := fmt.Sprintf("%d|%d|%d|%d", x+xd, y+yd, z+zd, w+wd)
							neighbor := universe[ncoord]
							neighbor.neighbors++
							universe[ncoord] = neighbor
						}
					}
				}
			}
		}
		for coords, c := range universe {
			if c.on && (c.neighbors < 2 || c.neighbors > 3) {
				delete(universe, coords)
			} else if !c.on && c.neighbors != 3 {
				delete(universe, coords)
			} else {
				universe[coords] = cube{on: true}
			}
		}
	}
	return len(universe)
}

func part1(universe map[string]cube) {
	fmt.Printf("The size of the 3d universe is %v\n", solve(universe, 3))
}

func part2(universe map[string]cube) {
	fmt.Printf("The size of the 4d universe is %v\n", solve(universe, 4))
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

func input() map[string]cube {
	ret := map[string]cube{}
	for y, line := range fileinput.ReadLines("input.txt") {
		for x, ch := range line {
			if ch == '#' {
				ret[fmt.Sprintf("%v|%v|0|0", x, y)] = cube{on: true}
			}
		}
	}
	return ret
}
