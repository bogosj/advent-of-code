package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() map[int]int {
	ret := map[int]int{}
	for _, line := range fileinput.ReadLines("input.txt") {
		f := strings.FieldsFunc(line, func(r rune) bool { return r == ',' || r == ' ' })
		f[0] = strings.ReplaceAll(f[0], ":", "")
		ret[intmath.Atoi(f[0])] = intmath.Atoi(f[1])
	}
	return ret
}

func traverse(delay int) (ret int) {
	for depth, size := range input() {
		period := size*2 - 2
		if (depth+delay)%period == 0 {
			ret += depth * size
		}
	}
	return
}

func part1() {
	fmt.Println("The severity of traversing the firewall is:", traverse(0))
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
