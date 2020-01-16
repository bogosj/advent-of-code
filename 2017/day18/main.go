package main

import (
	"fmt"
	"github.com/bogosj/advent-of-code/fileinput"
	"strconv"
	"strings"
	"time"
)

type computer struct {
	inst  [][]string
	reg   map[string]int
	pc, f int
}

func newComp() *computer {
	c := computer{reg: map[string]int{}}
	for _, line := range fileinput.ReadLines("input.txt") {
		c.inst = append(c.inst, strings.Fields(line))
	}
	return &c
}

func (c *computer) readVal(s string) int {
	if v, err := strconv.Atoi(s); err == nil {
		return v
	}
	return c.reg[s]
}

func (c *computer) runToFirstRCV() int {
	for {
		inst := c.inst[c.pc]
		c.pc++
		switch inst[0] {
		case "snd":
			c.f = c.reg[inst[1]]
		case "set":
			c.reg[inst[1]] = c.readVal(inst[2])
		case "add":
			c.reg[inst[1]] += c.readVal(inst[2])
		case "mul":
			c.reg[inst[1]] *= c.readVal(inst[2])
		case "mod":
			c.reg[inst[1]] %= c.readVal(inst[2])
		case "rcv":
			if c.readVal(inst[1]) != 0 {
				return c.f
			}
		case "jgz":
			if c.readVal(inst[1]) > 0 {
				c.pc += c.readVal(inst[2]) - 1
			}
		}
	}
}

func part1() {
	c := newComp()
	fmt.Println("The first fcv is", c.runToFirstRCV())
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
