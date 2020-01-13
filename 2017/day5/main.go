package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type instructions struct {
	i  []int
	pc int
}

func (i *instructions) run() (ret int) {
	for i.pc >= 0 && i.pc < len(i.i) {
		o := i.pc
		i.pc += i.i[i.pc]
		i.i[o]++
		ret++
	}
	return
}

func input() *instructions {
	i := instructions{}
	for _, line := range fileinput.ReadLines("input.txt") {
		i.i = append(i.i, intmath.Atoi(line))
	}
	return &i
}

func part1() {
	i := input()
	fmt.Printf("It takes %d steps to exit the program\n", i.run())
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
