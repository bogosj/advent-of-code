package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type computer struct {
	registers map[string]int
}

func newComp() *computer {
	c := computer{}
	c.registers = map[string]int{}
	return &c
}

func (c *computer) readVal(s string) int {
	switch s {
	case "a", "b", "c", "d":
		return c.registers[s]
	default:
		return intmath.Atoi(s)
	}
}

func (c *computer) runInstructions(p string) {
	inst := fileinput.ReadLines(p)
	pc := 0
	for pc < len(inst) {
		f := strings.Fields(inst[pc])
		pc++
		switch f[0] {
		case "cpy":
			c.registers[f[2]] = c.readVal(f[1])
		case "jnz":
			if c.readVal(f[1]) != 0 {
				pc--
				pc += intmath.Atoi(f[2])
			}
		case "inc":
			c.registers[f[1]]++
		case "dec":
			c.registers[f[1]]--
		}
	}
}

func part1() {
	c := newComp()
	c.runInstructions("input.txt")
	fmt.Println("The value in register a is:", c.registers["a"])
}

func part2() {
	c := newComp()
	c.registers["c"] = 1
	c.runInstructions("input.txt")
	fmt.Println("The value in register a with c=1 is:", c.registers["a"])
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
