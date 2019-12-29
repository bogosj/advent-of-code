package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/intmath"

	"github.com/bogosj/advent-of-code/fileinput"
)

func input() string {
	lines := fileinput.ReadLines("input.txt")
	return lines[0]
}

func part1() {
	m := map[intmath.Point]bool{}
	p := intmath.Point{}
	for _, r := range input() {
		m[p] = true
		switch r {
		case '^':
			p.Y--
		case '>':
			p.X++
		case '<':
			p.X--
		case 'v':
			p.Y++
		}
	}
	m[p] = true
	fmt.Println("Houses visited:", len(m))
}

func part2() {
	m := map[intmath.Point]bool{}
	p := []intmath.Point{intmath.Point{}, intmath.Point{}}
	for i, r := range input() {
		m[p[i%2]] = true
		switch r {
		case '^':
			p[i%2].Y--
		case '>':
			p[i%2].X++
		case '<':
			p[i%2].X--
		case 'v':
			p[i%2].Y++
		}
		m[p[i%2]] = true
	}
	fmt.Println("Houses visited:", len(m))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
