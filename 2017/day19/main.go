package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type point = intmath.Point

const (
	down = iota
	up
	left
	right
)

func input() map[point]rune {
	ret := map[point]rune{}
	for y, line := range fileinput.ReadLinesRaw("input.txt") {
		for x, r := range line {
			ret[point{X: x, Y: y}] = r
		}
	}
	return ret
}

func mapStart(m map[point]rune) (ret point) {
	for x := 0; ; x++ {
		ret = point{X: x, Y: 0}
		if m[ret] == '|' {
			return
		}
	}
}

func walkMap() (ret string) {
	var err error
	dir := down
	m := input()
	curr := mapStart(m)
	visited := map[point]bool{}
	for err == nil {
		visited[curr] = true
		curr = move(curr, dir)
		ret = checkForLetter(curr, m, ret)
		dir, err = maybeChangeDir(curr, m, visited, dir)
	}
	return
}

func maybeChangeDir(curr point, m map[point]rune, v map[point]bool, dir int) (newDir int, err error) {
	if m[curr] != '+' {
		return dir, nil
	}
	for _, n := range curr.Neighbors() {
		if v[n] {
			continue
		}
		if m[n] == '-' {
			if n.X > curr.X {
				return right, nil
			}
			return left, nil
		}
		if m[n] == '|' {
			if n.Y > curr.Y {
				return down, nil
			}
			return up, nil
		}
	}
	return dir, errors.New("end of the tube")
}

func checkForLetter(curr point, m map[point]rune, ret string) string {
	v := m[curr]
	if v >= 'A' && v <= 'Z' {
		ret += string(v)
	}
	return ret
}

func move(curr point, dir int) point {
	switch dir {
	case down:
		curr.Y++
	case up:
		curr.Y--
	case left:
		curr.X--
	case right:
		curr.X++
	}
	return curr
}

func part1() {
	fmt.Println("The letters found along the tube are:", walkMap())
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
