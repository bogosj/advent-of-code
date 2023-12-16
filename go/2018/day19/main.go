package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type computer struct {
	ipReg int
	ip    int
	reg   []int
}

func newComp() *computer {
	return &computer{
		reg: make([]int, 6),
	}
}

func vals(s []string) (ret []int) {
	for _, v := range s {
		ret = append(ret, intmath.Atoi(v))
	}
	return
}

func (c *computer) setIPReg(s string) {
	f := strings.Fields(s)
	c.ipReg = intmath.Atoi(f[1])
}

func (c *computer) runInst(f []string, v []int) {
	switch f[0] {
	case "addi":
		c.reg[v[2]] = c.reg[v[0]] + v[1]
	case "addr":
		c.reg[v[2]] = c.reg[v[0]] + c.reg[v[1]]
	case "eqrr":
		if c.reg[v[0]] == c.reg[v[1]] {
			c.reg[v[2]] = 1
		} else {
			c.reg[v[2]] = 0
		}
	case "gtrr":
		if c.reg[v[0]] > c.reg[v[1]] {
			c.reg[v[2]] = 1
		} else {
			c.reg[v[2]] = 0
		}
	case "muli":
		c.reg[v[2]] = c.reg[v[0]] * v[1]
	case "mulr":
		c.reg[v[2]] = c.reg[v[0]] * c.reg[v[1]]
	case "seti":
		c.reg[v[2]] = v[0]
	case "setr":
		c.reg[v[2]] = c.reg[v[0]]
	}
}

func (c *computer) run() {
	lines := fileinput.ReadLines("input.txt")
	c.setIPReg(lines[0])
	inst := lines[1:]
	for c.ip < len(inst) {
		f := strings.Fields(inst[c.ip])
		v := vals(f[1:])
		c.reg[c.ipReg] = c.ip
		c.runInst(f, v)
		c.ip = c.reg[c.ipReg]
		c.ip++
	}
}

func (c *computer) runFast() {
	lines := fileinput.ReadLines("input.txt")
	c.setIPReg(lines[0])
	inst := lines[1:]
	for c.ip < len(inst) {
		f := strings.Fields(inst[c.ip])
		v := vals(f[1:])
		c.reg[c.ipReg] = c.ip
		c.runInst(f, v)
		c.ip = c.reg[c.ipReg]
		c.ip++
		if c.reg[0] == 1 && c.reg[4] > 0 {
			c.reg[0] = intmath.Sum(intmath.Factors(c.reg[5])...)
			return
		}
	}
}

func part1() {
	c := newComp()
	c.runFast()
	fmt.Printf("The value in register 0 is: %d\n", c.reg[0])
}

func part2() {
	c := newComp()
	c.reg[0] = 1
	c.runFast()
	fmt.Printf("The value in register 0 is: %d\n", c.reg[0])
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
