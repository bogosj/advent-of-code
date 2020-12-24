package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

type hexPoint struct {
	x, y, z int
}

type directions struct {
	d string
}

func (d *directions) endPoint() (ret hexPoint) {
	for i := 0; i < len(d.d); i++ {
		switch d.d[i] {
		case 'e':
			ret.x++
			ret.y--
		case 'w':
			ret.x--
			ret.y++
		case 'n':
			i++
			ret.z--
			if d.d[i] == 'w' {
				ret.y++
			} else {
				ret.x++
			}
		case 's':
			i++
			ret.z++
			if d.d[i] == 'w' {
				ret.x--
			} else {
				ret.y--
			}
		}
	}
	return
}

func part1(in []directions) {
	m := map[hexPoint]bool{}
	for _, d := range in {
		p := d.endPoint()
		m[p] = !m[p]
	}
	count := 0
	for _, v := range m {
		if v {
			count++
		}
	}
	fmt.Printf("There are %d black tiles\n", count)
}

func part2(in []directions) {
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	in = input()
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []directions {
	ret := []directions{}

	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, directions{d: line})
	}

	return ret
}
