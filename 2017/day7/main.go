package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() (ret [][]string) {
	for _, lines := range fileinput.ReadLines("input.txt") {
		ret = append(ret, strings.Fields(lines))
	}
	return
}

type program struct {
	name   string
	weight int
	data   []string
	parent *program
}

func allPrograms() map[string]*program {
	ret := map[string]*program{}
	for _, line := range input() {
		p := program{name: line[0]}
		w := strings.ReplaceAll(line[1], "(", "")
		w = strings.ReplaceAll(w, ")", "")
		p.weight = intmath.Atoi(w)
		if len(line) > 3 {
			p.data = line[3:]
		}
		ret[p.name] = &p
	}
	return ret
}

func makeTree(m map[string]*program) {
	for _, v := range m {
		if v.data != nil {
			for _, name := range v.data {
				name = strings.ReplaceAll(name, ",", "")
				m[name].parent = v
			}
		}
	}
}

func part1() {
	a := allPrograms()
	makeTree(a)
	var root *program
	for _, v := range a {
		root = v
		for root.parent != nil {
			root = root.parent
		}
	}
	fmt.Println("The root program has the name:", root.name)
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
