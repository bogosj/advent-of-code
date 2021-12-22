package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type instruction struct {
	on      bool
	x, y, z []int
}

func buildInstructions(in []string) []instruction {
	ret := []instruction{}
	for _, line := range in {
		inst := instruction{}
		f := strings.Fields(line)
		inst.on = f[0] == "on"
		ranges := strings.Split(f[1], ",")
		xs := strings.Split(strings.Split(ranges[0], "=")[1], "..")
		inst.x = append(inst.x, intmath.Atoi(xs[0]))
		inst.x = append(inst.x, intmath.Atoi(xs[1]))
		ys := strings.Split(strings.Split(ranges[1], "=")[1], "..")
		inst.y = append(inst.y, intmath.Atoi(ys[0]))
		inst.y = append(inst.y, intmath.Atoi(ys[1]))
		zs := strings.Split(strings.Split(ranges[2], "=")[1], "..")
		inst.z = append(inst.z, intmath.Atoi(zs[0]))
		inst.z = append(inst.z, intmath.Atoi(zs[1]))
		ret = append(ret, inst)
	}
	return ret
}

func part1(in []string) {
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
