package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

type computer struct {
	inst     [][]string
	reg      map[string]int
	pc       int
	mulCount int
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

func (c *computer) run() {
	l := len(c.inst)
	for c.pc < l {
		inst := c.inst[c.pc]
		c.pc++
		switch inst[0] {
		case "set":
			c.reg[inst[1]] = c.readVal(inst[2])
		case "sub":
			c.reg[inst[1]] -= c.readVal(inst[2])
		case "mul":
			c.mulCount++
			c.reg[inst[1]] *= c.readVal(inst[2])
		case "jnz":
			if c.readVal(inst[1]) != 0 {
				c.pc--
				c.pc += c.readVal(inst[2])
			}
		}
	}
}

func part1() {
	c := newComp()
	c.run()
	fmt.Printf("The computer calls mul %d times\n", c.mulCount)
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
