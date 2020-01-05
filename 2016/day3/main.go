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
	i := 0
	data := input()
	for r := 0; r < len(data); r += 3 {
		for c := 0; c < 3; c++ {
			t := []int{data[r][c], data[r+1][c], data[r+2][c]}
			sort.Ints(t)
			if t[0]+t[1] > t[2] {
				i++
			}
		}
	}
	fmt.Println("There are", i, "valid triagles when using columns")
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
