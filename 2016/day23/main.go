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

func nextFiveInst(inst [][]string) (ret string) {
	for _, i:=range inst {
		ret+=i[0]
	}
	return
}

func (c *computer) runInstructions(p string) {
	var inst [][]string
	for _, line := range fileinput.ReadLines(p) {
		inst = append(inst, strings.Fields(line))
	}
	pc := 0
	for pc < len(inst) {
		f := inst[pc]
		pc++
		switch f[0] {
		case "cpy":
			c.registers[f[2]] = c.readVal(f[1])
			if nextFiveInst(inst[pc:pc+5]) == "incdecjnzdecjnz" {
				// MUL optimization
				c.registers[inst[pc][1]] += c.readVal(inst[pc+1][1]) * c.readVal(inst[pc+3][1])
				c.registers[inst[pc+1][1]] = 0
				c.registers[inst[pc+3][1]] = 0
				pc+=5
			}
		case "jnz":
			if c.readVal(f[1]) != 0 {
				pc--
				pc += c.readVal(f[2])
			}
		case "inc":
			c.registers[f[1]]++
		case "dec":
			c.registers[f[1]]--
		case "tgl":
			ni := pc + c.readVal(f[1]) - 1
			if ni >= 0 && ni < len(inst) {
				fi := inst[ni]
				if len(fi) == 2 {
					if fi[0] == "inc" {
						fi[0] = "dec"
					} else {
						fi[0] = "inc"
					}
				} else {
					if fi[0] == "jnz" {
						fi[0] = "cpy"
					} else {
						fi[1] = "jnz"
					}
				}
			}
		}
	}
}

func part1() {
	c := newComp()
	c.registers["a"] = 7
	c.runInstructions("input.txt")
	fmt.Println("The value in register a is:", c.registers["a"])
}

func part2() {
	c := newComp()
	c.registers["a"] = 12
	c.runInstructions("input.txt")
	fmt.Println("The value in register a is:", c.registers["a"])
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
