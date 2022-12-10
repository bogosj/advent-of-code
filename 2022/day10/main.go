package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func evalStrength(cycle, register int) int {
	switch cycle {
	case 20, 60, 100, 140, 180, 220:
		return register * cycle
	}
	return 0
}

func part1(in []string) {
	cycle := 1
	register := 1
	strength := 0
	for _, line := range in {
		inst := strings.Split(line, " ")[0]
		if inst == "noop" {
			strength += evalStrength(cycle, register)
			cycle++
		} else {
			strength += evalStrength(cycle, register)
			cycle++
			strength += evalStrength(cycle, register)
			cycle++
			register += intmath.Atoi(strings.Split(line, " ")[1])
		}
	}
	fmt.Printf("Signal strength: %d\n", strength)
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
