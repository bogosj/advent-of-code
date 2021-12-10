package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func part1(in []string) {
	i := 0
	count := 0
	for _, s := range in {
		if intmath.Atoi(s) > i {
			count++
		}
		i = intmath.Atoi(s)
	}
	fmt.Println("Part 1 answer:", count-1)
}

func part2(in []string) {
	i := 0
	count := 0
	for {
		i2 := intmath.Atoi(in[0]) + intmath.Atoi(in[1]) + intmath.Atoi(in[2])
		if i2 > i {
			count++
		}
		in = in[1:]
		if len(in) < 3 {
			break
		}
		i = i2
	}
	fmt.Println("Part 2 answer:", count-1)
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
