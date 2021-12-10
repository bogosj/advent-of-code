package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func part1(in []string) {
	pos := intmath.Point{}
	for _, inst := range in {
		f := strings.Fields(inst)
		switch f[0] {
		case "forward":
			pos.X += intmath.Atoi(f[1])
		case "up":
			pos.Y -= intmath.Atoi(f[1])
		case "down":
			pos.Y += intmath.Atoi(f[1])
		}
	}
	fmt.Println("Part 1 answer:", pos.X*pos.Y)
}

func part2(in []string) {
	pos := intmath.Point{}
	aim := 0
	for _, inst := range in {
		f := strings.Fields(inst)
		switch f[0] {
		case "forward":
			pos.X += intmath.Atoi(f[1])
			pos.Y += aim * intmath.Atoi(f[1])
		case "up":
			aim -= intmath.Atoi(f[1])
		case "down":
			aim += intmath.Atoi(f[1])
		}
	}
	fmt.Println("Part 1 answer:", pos.X*pos.Y)
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
