package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

var (
	memReg = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
)

type computer struct {
	inst []string
	mem  map[int]int
	mask string
}

func (c *computer) setMem(line string) {
	match := memReg.FindStringSubmatch(line)
	addr := intmath.Atoi(match[1])
	val := intmath.Atoi(match[2])
	for i, bit := 0, 1<<35; i < len(c.mask); i, bit = i+1, bit>>1 {
		switch c.mask[i] {
		case '0':
			val = val &^ bit
		case '1':
			val = val | bit
		}
	}
	c.mem[addr] = val
}

func (c *computer) init() {
	for _, line := range c.inst {
		if strings.HasPrefix(line, "mask") {
			c.mask = strings.TrimPrefix(line, "mask = ")
		} else {
			c.setMem(line)
		}
	}
}

func part1(in *computer) {
	in.init()
	a := 0
	for _, v := range in.mem {
		a += v
	}
	fmt.Printf("The sum of values is %v\n", a)
}

func part2(in *computer) {
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() *computer {
	return &computer{
		inst: fileinput.ReadLines("input.txt"),
		mem:  map[int]int{},
	}
}
