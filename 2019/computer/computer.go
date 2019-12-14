package computer

import (
	"errors"
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

type Computer struct {
	prog   map[int]int
	pc, rc int
}

func New(prog map[int]int) *Computer {
	c := Computer{prog: prog}
	return &c
}

func (c *Computer) Hack(addr, val int) {
	c.prog[addr] = val
}

func (c *Computer) Compute(in int) (int, error) {
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
			c.prog[idx] = in
			c.pc += 2
			//ic++
		case 4: // Output
			c.pc += 2
			return vals[0], nil
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
			return 0, errors.New("halt")
		}
	}
}
