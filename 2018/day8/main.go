package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

var (
	simple = []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}
)

func input() (ret []int) {
	lines := fileinput.ReadLines("input.txt")
	for _, f := range strings.Fields(lines[0]) {
		ret = append(ret, intmath.Atoi(f))
	}
	return
}

type node struct {
	children []*node
	metadata []int
}

func (n *node) value() int {
	if len(n.children) == 0 {
		return intmath.Sum(n.metadata...)
	}

	var ret int
	for _, v := range n.metadata {
		if v <= len(n.children) {
			ret += n.children[v-1].value()
		}
	}
	return ret
}

func tree(in []int) (*node, []int) {
	n := node{}
	children := in[0]
	metaLen := in[1]
	var nc *node
	in = in[2:]
	for i := 0; i < children; i++ {
		nc, in = tree(in)
		n.children = append(n.children, nc)
	}
	for i := 0; i < metaLen; i++ {
		n.metadata = append(n.metadata, in[0])
		in = in[1:]
	}
	return &n, in
}

func checksum(n *node) (ret int) {
	ret += intmath.Sum(n.metadata...)
	for _, c := range n.children {
		ret += checksum(c)
	}
	return
}

func part1() {
	t, _ := tree(input())
	fmt.Println("The checksum of the tree is:", checksum(t))
}

func part2() {
	t, _ := tree(input())
	fmt.Println("The value of the root node is:", t.value())
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
