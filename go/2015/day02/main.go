package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/intmath"

	"github.com/bogosj/advent-of-code/fileinput"
)

func input() (ret [][]int) {
	lines := fileinput.ReadLines("input.txt")
	for _, line := range lines {
		f := strings.FieldsFunc(line, func(r rune) bool { return r == 'x' })
		row := []int{}
		for _, num := range f {
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			row = append(row, i)
		}
		ret = append(ret, row)
	}
	return
}

func part1() {
	sum := 0
	for _, dim := range input() {
		a := dim[0] * dim[1] * 2
		b := dim[1] * dim[2] * 2
		c := dim[0] * dim[2] * 2
		d := intmath.Min(a, b, c) / 2
		sum += a + b + c + d
	}
	fmt.Println("Need:", sum)
}

func part2() {
	sum := 0
	for _, dim := range input() {
		sum += (dim[0] * dim[1] * dim[2])
		max := intmath.Max(dim...)
		sum += 2*(dim[0]+dim[1]+dim[2]) - 2*max
	}
	fmt.Println("Ribbon:", sum)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
