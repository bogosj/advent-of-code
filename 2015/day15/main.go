package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input(p string) map[string]map[string]int {
	ret := map[string]map[string]int{}
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		f := strings.FieldsFunc(line, func(r rune) bool {
			return r == ':'
		})
		f2 := strings.FieldsFunc(f[1], func(r rune) bool {
			return r == ','
		})
		ret[f[0]] = map[string]int{}
		for _, attr := range f2 {
			f3 := strings.Fields(attr)
			ret[f[0]][f3[0]] = intmath.Atoi(f3[1])
		}
	}
	return ret
}

func part1() {
	fmt.Println(input("input.txt"))
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
