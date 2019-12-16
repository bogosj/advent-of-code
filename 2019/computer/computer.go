package computer

import (
	"fmt"
	"strconv"
	"strings"

	"jamesbogosian.com/advent-of-code/2019/fileinput"
)

type opCode struct {
	code  int
	modes []int
}

func (o *opCode) ReadVals() bool {
	return o.code != 99
}

func parseOpCode(i int) opCode {
	ret := opCode{code: i % 100}
	i /= 100
	for i != 0 {
		ret.modes = append(ret.modes, i%10)
		i /= 10
	}
	for len(ret.modes) < 3 {
		ret.modes = append(ret.modes, 0)
	}
	return ret
}

// Computer represents a IntCode computer.
type Computer struct {
	origProg map[int]int
	prog     map[int]int
	pc, rc   int
	Halted   bool
}

// New creates a new Computer from an input text file.
func New(path string) *Computer {
	c := Computer{origProg: input(path)}
	c.Reset()
	return &c
}

// Reset restores the computer's program memory to its original state.
func (c *Computer) Reset() {
	c.prog = map[int]int{}
	for k, v := range c.origProg {
		c.prog[k] = v
	}
}

// Hack allows a user to alter the memory of a specified address.
func (c *Computer) Hack(addr, val int) {
	c.prog[addr] = val
}

// Compute runs the computation of the program for a set of inputs.
func (c *Computer) Compute(in ...int) int {
	for {
		op := parseOpCode(c.prog[c.pc])
		vals := []int{}
		if op.ReadVals() {
			for i, mode := range op.modes {
				switch mode {
				case 0:
					vals = append(vals, c.prog[c.prog[c.pc+i+1]])
				case 1:
					vals = append(vals, c.prog[c.pc+i+1])
				default:
					vals = append(vals, c.prog[c.prog[c.pc+i+1]+c.rc])
				}
			}
		}
		switch op.code {
		case 1: // Add
			idx := c.prog[c.pc+3]
			if op.modes[2] == 2 {
				idx += c.rc
			}
			c.prog[idx] = vals[0] + vals[1]
			c.pc += 4
		case 2: // Multiply
			idx := c.prog[c.pc+3]
			if op.modes[2] == 2 {
				idx += c.rc
			}
			c.prog[idx] = vals[0] * vals[1]
			c.pc += 4
		case 3: // Store
			idx := c.prog[c.pc+1]
			if op.modes[0] == 2 {
				idx += c.rc
			}
			c.prog[idx] = in[0]
			in = in[1:]
			c.pc += 2
		case 4: // Output
			c.pc += 2
			return vals[0]
		case 5:
			if vals[0] != 0 {
				c.pc = vals[1]
			} else {
				c.pc += 3
			}
		case 6:
			if vals[0] == 0 {
				c.pc = vals[1]
			} else {
				c.pc += 3
			}
		case 7:
			idx := c.prog[c.pc+3]
			if op.modes[2] == 2 {
				idx += c.rc
			}
			if vals[0] < vals[1] {
				c.prog[idx] = 1
			} else {
				c.prog[idx] = 0
			}
			c.pc += 4
		case 8:
			idx := c.prog[c.pc+3]
			if op.modes[2] == 2 {
				idx += c.rc
			}
			if vals[0] == vals[1] {
				c.prog[idx] = 1
			} else {
				c.prog[idx] = 0
			}
			c.pc += 4
		case 9:
			c.rc += vals[0]
			c.pc += 2
		case 99:
			c.Halted = true
			return 0
		}
	}
}

func input(n string) map[int]int {
	ret := map[int]int{}
	lines := fileinput.ReadLines(n)
	for i, v := range strings.Split(lines[0], ",") {
		iv, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		ret[i] = iv
	}
	return ret
}
