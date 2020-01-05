package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() (ret [][]int) {
	lines := fileinput.ReadLines("input.txt")
	for _, line := range lines {
		f := strings.Fields(line)
		var row []int
		for _, v := range f {
			row = append(row, intmath.Atoi(v))
		}
		ret = append(ret, row)
	}
	return
}

func part1() {
	i := 0
	for _, row := range input() {
		sort.Ints(row)
		if row[0]+row[1] > row[2] {
			i++
		}
	}
	fmt.Println("There are", i, "valid triangles")
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
