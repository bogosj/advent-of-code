package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type point = intmath.Point

type pixel struct {
	p, v point
}

func (p *pixel) advance() {
	p.p.X += p.v.X
	p.p.Y += p.v.Y
}

func atoi(s string) int {
	return intmath.Atoi(strings.TrimSpace(s))
}

func input() (ret []*pixel) {
	for _, line := range fileinput.ReadLines("input.txt") {
		p := pixel{}
		p.p = point{
			X: atoi(line[10:16]),
			Y: atoi(line[18:24]),
		}
		p.v = point{
			X: atoi(line[36:38]),
			Y: atoi(line[40:42]),
		}
		ret = append(ret, &p)
	}
	return
}

func minMax(pixels []*pixel) (minX, minY, maxX, maxY int) {
	minX = 9999999
	minY = 9999999
	maxX = -9999999
	maxY = -9999999
	for _, p := range pixels {
		minX = intmath.Min(minX, p.p.X)
		minY = intmath.Min(minY, p.p.Y)
		maxX = intmath.Max(maxX, p.p.X)
		maxY = intmath.Max(maxY, p.p.Y)
	}
	return
}

func advance(pixels []*pixel) {
	for _, p := range pixels {
		p.advance()
	}
}

func printPixels(pixels []*pixel, seconds int) {
	m := map[point]bool{}
	for _, p := range pixels {
		m[p.p] = true
	}
	_, minY, maxX, maxY := minMax(pixels)
	if maxY-minY < 30 {
		for y := 0; y <= maxY; y++ {
			for x := 0; x <= maxX; x++ {
				if m[point{X: x, Y: y}] {
					fmt.Printf("#")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
		fmt.Printf("============= %d\n", seconds)
	}
}

func part1() {
	pixels := input()
	for i := 0; i < 14000; i++ {
		advance(pixels)
		printPixels(pixels, i+1)
	}
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
