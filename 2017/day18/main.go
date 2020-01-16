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

func (c *computer) run(in <-chan int) <-chan int {
	out := make(chan int, 5000)
	go func() {
		for {
			inst := c.inst[c.pc]
			c.pc++
			switch inst[0] {
			case "snd":
				out <- c.reg[inst[1]]
			case "set":
				c.reg[inst[1]] = c.readVal(inst[2])
			case "add":
				c.reg[inst[1]] += c.readVal(inst[2])
			case "mul":
				c.reg[inst[1]] *= c.readVal(inst[2])
			case "mod":
				c.reg[inst[1]] %= c.readVal(inst[2])
			case "rcv":
				c.reg[inst[1]] = <-in
			case "jgz":
				if c.readVal(inst[1]) > 0 {
					c.pc += c.readVal(inst[2]) - 1
				}
			}
		}
	}()
	return out
}

func part1() {
	c := newComp()
	fmt.Println("The first fcv is", c.runToFirstRCV())
}

func part2() {
	c1 := newComp()
	c2 := newComp()
	c1.reg["p"] = 0
	c2.reg["p"] = 1
	in1 := make(chan int)
	in2 := make(chan int)
	o1 := c1.run(in1)
	o2 := c2.run(in2)
	var c, sleep int
	for {
		select {
		case m := <-o1:
			sleep = 0
			in2 <- m
		case m := <-o2:
			sleep = 0
			c++
			in1 <- m
		default:
			time.Sleep(time.Second)
			sleep++
		}
		if sleep >= 5 {
			fmt.Printf("Computers have not responded for 5s, C2 sent %d messages\n", c)
			return
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
