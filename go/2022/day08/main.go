package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func part1(in map[intmath.Point]int) {
	visible := map[intmath.Point]bool{}
	for y := 0; y < 99; y++ {
		for x := 0; x < 99; x++ {
			markPoint(x, y, in, visible)
		}
	}
	fmt.Printf("There are %d visible trees.\n", len(visible))
}

func markPoint(x int, y int, in map[intmath.Point]int, visible map[intmath.Point]bool) {
	currPoint := intmath.Point{X: x, Y: y}
	height := in[currPoint]
	// Look up, down, left, right. If pass edge, set visible=true and return
	// Up
	for y2 := y - 1; y2 >= -1; y2-- {
		if y2 == -1 {
			visible[currPoint] = true
			return
		}
		if in[intmath.Point{X: x, Y: y2}] >= height {
			break
		}
	}
	// Down
	for y2 := y + 1; y2 <= 99; y2++ {
		if y2 == 99 {
			visible[currPoint] = true
			return
		}
		if in[intmath.Point{X: x, Y: y2}] >= height {
			break
		}
	}
	// Left
	for x2 := x - 1; x2 >= -1; x2-- {
		if x2 == -1 {
			visible[currPoint] = true
			return
		}
		if in[intmath.Point{X: x2, Y: y}] >= height {
			break
		}
	}
	// Right
	for x2 := x + 1; x2 <= 99; x2++ {
		if x2 == 99 {
			visible[currPoint] = true
			return
		}
		if in[intmath.Point{X: x2, Y: y}] >= height {
			break
		}
	}
}

func calculateScenicScore(x, y int, in map[intmath.Point]int) int {
	ret := 1
	tree := intmath.Point{X: x, Y: y}
	height := in[tree]

	// Up
	for _, dir := range "UDLR" {
		neighbor := tree.Neighbor(dir)
		count := 0
		for {
			nextTree, ok := in[neighbor]
			if !ok {
				break
			}
			count++
			if nextTree >= height {
				break
			}
			neighbor = neighbor.Neighbor(dir)
		}
		ret *= count
	}
	return ret
}

func part2(in map[intmath.Point]int) {
	scores := []int{}
	for y := 0; y < 99; y++ {
		for x := 0; x < 99; x++ {
			scores = append(scores, calculateScenicScore(x, y, in))
		}
	}
	fmt.Printf("The max scenic score is %d\n", intmath.Max(scores...))
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

func input() map[intmath.Point]int {
	ret := map[intmath.Point]int{}
	for y, line := range fileinput.ReadLines("input.txt") {
		for x, char := range line {
			ret[intmath.Point{X: x, Y: y}] = int(char - 48)
		}
	}
	return ret
}
