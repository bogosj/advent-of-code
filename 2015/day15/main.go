package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func input(p string) map[string]map[string]int {
	ret := map[string]map[string]int{}
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		f := strings.FieldsFunc(line, func(r rune) bool {
			return r == ':'
		})
		for _, l := range f {
			fmt.Println(l)
		}
	}
	return ret
}

func part1() {
	input("input.txt")
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
