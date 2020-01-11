package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() (ret []node) {
	lines := fileinput.ReadLines("input.txt")
	for _, line := range lines[2:] {
		f := strings.Fields(line)
		n := node{}
		n.name = f[0]
		n.size = intmath.Atoi(strings.ReplaceAll(f[1], "T", ""))
		n.used = intmath.Atoi(strings.ReplaceAll(f[2], "T", ""))
		n.avail = intmath.Atoi(strings.ReplaceAll(f[3], "T", ""))
		n.setLoc()
		ret = append(ret, n)
	}
	return
}

func nodeArray() map[intmath.Point]node {
	ret := map[intmath.Point]node{}
	for _, n := range input() {
		ret[n.loc] = n
	}
	return ret
}

type node struct {
	name              string
	loc               intmath.Point
	size, used, avail int
}

func (n *node) String() string {
	return fmt.Sprintf("%s: %dT | %dT | %dT", n.name, n.size, n.used, n.avail)
}

func (n *node) setLoc() {
	f := strings.FieldsFunc(n.name, func(r rune) bool { return r == '-' })
	n.loc = intmath.Point{}
	n.loc.X = intmath.Atoi(strings.ReplaceAll(f[1], "x", ""))
	n.loc.Y = intmath.Atoi(strings.ReplaceAll(f[2], "y", ""))
}

func part1() {
	in := input()
	pairs := 0
	for i := 0; i < len(in); i++ {
		n1 := in[i]
		if n1.used == 0 {
			continue
		}
		for j := 0; j < len(in); j++ {
			n2 := in[j]
			if n1 != n2 && n1.used <= n2.avail {
				pairs++
			}
		}
	}
	fmt.Println("There are", pairs, "viable node pairs.")
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
