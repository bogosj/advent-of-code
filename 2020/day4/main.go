package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/intmath"

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

func (p *passport) validEcl() bool {
	for _, v := range strings.Split("amb blu brn gry grn hzl oth", " ") {
		if p.data["ecl"] == v {
			return true
		}
	}
	return false
}

func (p *passport) isValid2() bool {
	if !p.isValid1() {
		return false
	}

	byr := intmath.Atoi(p.data["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}

	iyr := intmath.Atoi(p.data["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr := intmath.Atoi(p.data["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	hgt := p.data["hgt"]
	if strings.HasSuffix(hgt, "cm") {
		h := intmath.Atoi(strings.TrimSuffix(hgt, "cm"))
		if h < 150 || h > 193 {
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		h := intmath.Atoi(strings.TrimSuffix(hgt, "in"))
		if h < 59 || h > 76 {
			return false
		}
	} else {
		return false
	}

	if m, _ := regexp.MatchString(`^#([0-9a-f]){6}$`, p.data["hcl"]); !m {
		return false
	}

	if m, _ := regexp.MatchString(`^([0-9]){9}$`, p.data["pid"]); !m {
		return false
	}

	return p.validEcl()
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
	c := 0
	for _, p := range in {
		if p.isValid2() {
			c++
		}
	}
	fmt.Printf("There are %v really valid passports.\n", c)
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
