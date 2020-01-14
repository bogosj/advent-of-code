package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() (ret [][]string) {
	for _, lines := range fileinput.ReadLines("input.txt") {
		ret = append(ret, strings.Fields(lines))
	}
	return
}

type computer struct {
	inst      [][]string
	registers map[string]int
}

func newComp() *computer {
	c := computer{}
	c.inst = input()
	c.registers = map[string]int{}
	return &c
}

func (c *computer) test(i []string) bool {
	reg := i[4]
	v := intmath.Atoi(i[6])
	switch i[5] {
	case "<=":
		return c.registers[reg] <= v
	case ">=":
		return c.registers[reg] >= v
	case "==":
		return c.registers[reg] == v
	case "!=":
		return c.registers[reg] != v
	case ">":
		return c.registers[reg] > v
	case "<":
		return c.registers[reg] < v
	}
	return false
}

func (c *computer) execute(i []string) {
	reg := i[0]
	v := intmath.Atoi(i[2])
	if i[1] == "inc" {
		c.registers[reg] += v
	} else {
		c.registers[reg] -= v
	}
}

func (c *computer) run() (reg string, val int) {
	for _, i := range c.inst {
		if c.test(i) {
			c.execute(i)
		}
		r, v := c.largestValue()
		if v > val {
			reg = r
			val = v
		}
	}
	return
}

func (c *computer) largestValue() (reg string, val int) {
	for k, v := range c.registers {
		if v > val {
			val = v
			reg = k
		}
	}
	return
}

func part1() {
	c := newComp()
	c.run()
	r, v := c.largestValue()
	fmt.Printf("The largest value is in register %s: %d\n", r, v)
}

func part2() {
	c := newComp()
	r, v := c.run()
	fmt.Printf("The largest value during execution is in register %s: %d\n", r, v)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
