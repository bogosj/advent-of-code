package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type password struct {
	min, max int
	char     rune
	pass     string
}

func (p *password) isValid() bool {
	count := 0
	for _, c := range p.pass {
		if c == p.char {
			count++
		}
	}
	return count >= p.min && count <= p.max
}

func part1() {
	count := 0
	for _, p := range input() {
		if p.isValid() {
			count++
		}
	}
	fmt.Printf("There are %v valid passwords\n", count)
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []password {
	var ret []password
	lines := fileinput.ReadLines("input.txt")
	for _, v := range lines {
		s := strings.Split(v, " ")
		r := strings.Split(s[0], "-")
		p := password{
			min:  intmath.Atoi(r[0]),
			max:  intmath.Atoi(r[1]),
			char: rune(s[1][0]),
			pass: s[2],
		}
		ret = append(ret, p)
	}
	return ret
}
