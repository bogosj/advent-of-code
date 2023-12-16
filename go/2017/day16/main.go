package main

import (
	"fmt"
	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
	"strings"
	"time"
)

func input() (ret []string) {
	lines := fileinput.ReadLines("input.txt")
	return strings.FieldsFunc(lines[0], func(r rune) bool { return r == ',' })
}

func programs() (ret []rune) {
	r := 'a'
	for r <= 'p' {
		ret = append(ret, r)
		r++
	}
	return
}

func spin(p []rune, n int) []rune {
	return append(p[len(p)-n:], p[:len(p)-n]...)
}

func exchange(p []rune, a, b int) {
	p[a], p[b] = p[b], p[a]
}

func partner(p []rune, a, b rune) {
	var ai, bi int
	for i, r := range p {
		if r == a {
			ai = i
		} else if r == b {
			bi = i
		}
	}
	exchange(p, ai, bi)
}

func split(s string) (a, b string) {
	f := strings.FieldsFunc(s, func(r rune) bool { return r == '/' })
	return f[0], f[1]
}

func intSplit(s string) (a, b int) {
	as, bs := split(s)
	return intmath.Atoi(as), intmath.Atoi(bs)
}

func dance(p []rune) []rune {
	for _, inst := range input() {
		rest := inst[1:]
		switch inst[0] {
		case 's':
			p = spin(p, intmath.Atoi(rest))
		case 'x':
			a, b := intSplit(rest)
			exchange(p, a, b)
		case 'p':
			a, b := split(rest)
			partner(p, rune(a[0]), rune(b[0]))
		}
	}
	return p
}

func part1() {
	p := programs()
	p = dance(p)
	fmt.Println(string(p))
}

func part2() {
	seen := map[int]string{}
	p := programs()
	for i := 0; i < 48; i++ {
		seen[i] = string(p)
		p = dance(p)
	}
	fmt.Println("After a billion runs, they stand:", seen[1000000000%48])
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
