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

func (c *computer) runInst(f string, v []int) {
	switch f {
	case "addi":
		c.reg[v[2]] = c.reg[v[0]] + v[1]
	case "addr":
		c.reg[v[2]] = c.reg[v[0]] + c.reg[v[1]]
	case "bani":
		c.reg[v[2]] = c.reg[v[0]] & v[1]
	case "bori":
		c.reg[v[2]] = c.reg[v[0]] | v[1]
	case "eqri":
		if c.reg[v[0]] == v[1] {
			c.reg[v[2]] = 1
		} else {
			c.reg[v[2]] = 0
		}
	case "eqrr":
		if c.reg[v[0]] == c.reg[v[1]] {
			c.reg[v[2]] = 1
		} else {
			c.reg[v[2]] = 0
		}
	case "gtir":
		if v[0] > c.reg[v[1]] {
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
	case "seti":
		c.reg[v[2]] = v[0]
	case "setr":
		c.reg[v[2]] = c.reg[v[0]]
	}
}

func (c *computer) run(findOne bool) <-chan int {
	out := make(chan int, 10000)
	go func() {
		defer close(out)
		lines := fileinput.ReadLines("input.txt")
		c.setIPReg(lines[0])
		inst := lines[1:]
		var ops []string
		var vs [][]int
		for _, i := range inst {
			f := strings.Fields(i)
			ops = append(ops, f[0])
			vs = append(vs, vals(f[1:]))
		}
		for c.ip < len(inst) {
			if c.ip == 28 {
				out <- c.reg[4]
				if findOne {
					return
				}
			}
			c.reg[c.ipReg] = c.ip
			c.runInst(ops[c.ip], vs[c.ip])
			c.ip = c.reg[c.ipReg]
			c.ip++
		}
	}()
	return out
}

func part1() {
	c := newComp()
	reg := c.run(true)
	for {
		select {
		case val := <-reg:
			fmt.Println("The first halt is at:", val)
			return
		default:
			time.Sleep(time.Second)
		}
	}
}

func part2() {
	seen := map[int]bool{}
	c := newComp()
	reg := c.run(false)
	var prev int
OUTER:
	for {
		select {
		case val := <-reg:
			if seen[val] == true {
				fmt.Println("The lowest value to halt after the most instructions is:", prev)
				break OUTER
			}
			prev = val
			seen[val] = true
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
