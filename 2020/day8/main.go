package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type instruction struct {
	code string
	val  int
}

type console struct {
	pc, acc  int
	inst     []instruction
	executed map[int]bool
}

func (c *console) executeUntilDupe() {
	for !c.executed[c.pc] {
		c.executed[c.pc] = true
		i := c.inst[c.pc]
		switch i.code {
		case "nop":
			c.pc++
		case "acc":
			c.acc += i.val
			c.pc++
		case "jmp":
			c.pc += i.val
		}
	}
}

func newConsole(in []string) console {
	c := console{}
	for _, line := range in {
		f := strings.Split(line, " ")
		c.inst = append(c.inst, instruction{code: f[0], val: intmath.Atoi(f[1])})
	}
	c.executed = map[int]bool{}
	return c
}

func part1(in []string) {
	c := newConsole(in)
	c.executeUntilDupe()
	fmt.Printf("The val of acc before a dupe instruction is: %v\n", c.acc)
}

func part2(in []string) {
}

func main() {
	in := fileinput.ReadLines("input.txt")
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}
