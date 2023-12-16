package computer

import (
	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
	"strings"
)

type Computer struct {
	inst [][]string
	pc   int
	Reg  map[string]int
}

func New(p string) *Computer {
	c := Computer{}
	c.load(p)
	c.Reg = map[string]int{}
	return &c
}

func (c *Computer) load(p string) {
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		c.inst = append(c.inst, strings.Fields(line))
	}
}

func offset(s string) int {
	ret := intmath.Atoi(s[1:])
	if s[0] == '-' {
		ret *= -1
	}
	return ret
}

func (c *Computer) Run() {
	for c.pc < len(c.inst) {
		inst := c.inst[c.pc]
		r := inst[1]
		switch inst[0] {
		case "hlf":
			c.Reg[r] /= 2
			c.pc++
		case "tpl":
			c.Reg[r] *= 3
			c.pc++
		case "inc":
			c.Reg[r]++
			c.pc++
		case "jmp":
			c.pc += offset(r)
		case "jie":
			reg := string(r[0])
			if c.Reg[reg]%2 == 0 {
				c.pc += offset(inst[2])
			} else {
				c.pc++
			}
		case "jio":
			reg := string(r[0])
			if c.Reg[reg] == 1 {
				c.pc += offset(inst[2])
			} else {
				c.pc++
			}
		}
	}
}
