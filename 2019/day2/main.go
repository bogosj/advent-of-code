package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bogosj/advent-of-code/2019/fileinput"
)

func compute(in []int, n, v int) int {
	p := in
	p[1] = n
	p[2] = v
	s := 0
	for {
		f := p[s : s+4]
		op := f[0]
		if op == 1 {
			p[f[3]] = p[f[1]] + p[f[2]]
		} else if op == 2 {
			p[f[3]] = p[f[1]] * p[f[2]]
		} else if op == 99 {
			break
		}
		s += 4
	}
	return p[0]
}

func main() {
	fmt.Printf("result: %v\n", compute(input(), 12, 2))

	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			if compute(input(), n, v) == 19690720 {
				fmt.Printf("n=%v v=%v, a=%v\n", n, v, 100*n+v)
			}
		}
	}
}

func input() []int {
	var ret []int
	lines := fileinput.ReadLines("input.txt")
	for _, v := range strings.Split(lines[0], ",") {
		iv, _ := strconv.Atoi(v)
		ret = append(ret, iv)
	}
	return ret
}
