package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

type passport struct {
	data map[string]string
}

func (p *passport) isValid1() bool {
	for _, k := range strings.Split("byr iyr eyr hgt hcl ecl pid", " ") {
		if _, ok := p.data[k]; !ok {
			return false
		}
	}
	return true
}

func newPassport() passport {
	p := passport{}
	p.data = map[string]string{}
	return p
}

func part1(in []passport) {
	c := 0
	for _, p := range in {
		if p.isValid1() {
			c++
		}
	}
	fmt.Printf("There are %v valid passports.\n", c)
}

func part2(in []passport) {
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

func input() []passport {
	ret := []passport{}
	p := newPassport()
	for _, line := range fileinput.ReadLines("input.txt") {
		if strings.TrimSpace(line) == "" {
			ret = append(ret, p)
			p = newPassport()
		} else {
			d := strings.Split(line, " ")
			for _, item := range d {
				v := strings.Split(item, ":")
				p.data[v[0]] = v[1]
			}
		}
	}
	ret = append(ret, p)

	return ret
}