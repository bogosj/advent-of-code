package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type sample struct {
	before, op, after []int
}

func (s sample) test(f func(a, b, c int, reg []int)) bool {
	b := append([]int(nil), s.before...)
	a := append([]int(nil), s.after...)
	f(s.op[1], s.op[2], s.op[3], b)
	return reflect.DeepEqual(a, b)
}

func sampleInput() (ret []sample) {
	var prevLine string
	lines := fileinput.ReadLines("input.txt")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" && prevLine == "" {
			return
		}
		if strings.HasPrefix(line, "Before") {
			ns := sample{}
			ss := strings.Split(line, "[")[1]
			ss = strings.Split(ss, "]")[0]
			for _, v := range strings.Split(ss, ", ") {
				ns.before = append(ns.before, intmath.Atoi(v))
			}
			i++
			line = lines[i]
			for _, v := range strings.Split(line, " ") {
				ns.op = append(ns.op, intmath.Atoi(v))
			}
			i++
			line = lines[i]
			ss = strings.Split(line, "[")[1]
			ss = strings.Split(ss, "]")[0]
			for _, v := range strings.Split(ss, ", ") {
				ns.after = append(ns.after, intmath.Atoi(v))
			}
			ret = append(ret, ns)
		}

		prevLine = line
	}
	return
}

func programInput() (ret [][]int) {
	lines := fileinput.ReadLines("input.txt")
	var i int
	for !(lines[i] == "" && lines[i+1] == "" && lines[i+2] == "") {
		i++
	}
	for ; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		var n []int
		for _, v := range strings.Fields(line) {
			n = append(n, intmath.Atoi(v))
		}
		ret = append(ret, n)
	}
	return
}

func addr(a, b, c int, reg []int) {
	reg[c] = reg[a] + reg[b]
}
func addi(a, b, c int, reg []int) {
	reg[c] = reg[a] + b
}
func mulr(a, b, c int, reg []int) {
	reg[c] = reg[a] * reg[b]
}
func muli(a, b, c int, reg []int) {
	reg[c] = reg[a] * b
}
func banr(a, b, c int, reg []int) {
	reg[c] = reg[a] & reg[b]
}
func bani(a, b, c int, reg []int) {
	reg[c] = reg[a] & b
}
func borr(a, b, c int, reg []int) {
	reg[c] = reg[a] | reg[b]
}
func bori(a, b, c int, reg []int) {
	reg[c] = reg[a] | b
}
func setr(a, b, c int, reg []int) {
	reg[c] = reg[a]
}
func seti(a, b, c int, reg []int) {
	reg[c] = a
}
func gtir(a, b, c int, reg []int) {
	if a > reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}
func gtri(a, b, c int, reg []int) {
	if reg[a] > b {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}
func gtrr(a, b, c int, reg []int) {
	if reg[a] > reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}
func eqir(a, b, c int, reg []int) {
	if a == reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}
func eqri(a, b, c int, reg []int) {
	if reg[a] == b {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}
func eqrr(a, b, c int, reg []int) {
	if reg[a] == reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func part1() {
	ops := []func(a, b, c int, reg []int){addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
	var threes int
	samples := sampleInput()
	for _, s := range samples {
		var i int
		for _, f := range ops {
			if s.test(f) {
				i++
			}
		}
		if i >= 3 {
			threes++
		}
	}
	fmt.Printf("Of %d samples there were %d that behave like three or more opcodes.\n", len(samples), threes)
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
