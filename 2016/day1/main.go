package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

const (
	north = iota
	east
	south
	west
)

func input() (ret []string) {
	lines := fileinput.ReadLines("input.txt")
	fs := strings.FieldsFunc(lines[0], func(r rune) bool {
		return r == ' ' || r == ','
	})
	for _, f := range fs {
		ret = append(ret, string(f[0]), f[1:])
	}
	return
}

func walk(inst []string) (x, y int) {
	dir := north
	for _, i := range inst {
		dir += 4
		switch i {
		case "R":
			dir++
		case "L":
			dir--
		default:
			d := intmath.Atoi(i)
			switch dir % 4 {
			case north:
				y -= d
			case south:
				y += d
			case east:
				x += d
			case west:
				x -= d
			}
		}
	}
	return
}

func part1() {
	fmt.Println(walk(input()))
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
