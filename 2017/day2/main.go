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
		var r []int
		for _, v := range strings.Fields(line) {
			r = append(r, intmath.Atoi(v))
		}
		ret = append(ret, r)
	}
	return
}

func checksum(i []int) int {
	return intmath.Max(i...) - intmath.Min(i...)
}

func checksum2(in []int) int {
	sort.Ints(in)
	for i := len(in) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if in[i]%in[j] == 0 {
				return in[i] / in[j]
			}
		}
	}
	return -1
}

func part1() {
	var sum int
	for _, row := range input() {
		sum += checksum(row)
	}
	fmt.Println("The checksum of the spreadsheet is:", sum)
}

func part2() {
	var sum int
	for _, row := range input() {
		sum += checksum2(row)
	}
	fmt.Println("The part 2 checksum of the spreadsheet is:", sum)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
