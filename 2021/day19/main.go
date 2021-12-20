package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type scanner struct {
	points [][]int
	id     int
}

func (s *scanner) addPoints(line string) {
	sps := strings.Split(line, ",")
	p := make([]int, 3)
	for i, sp := range sps {
		p[i] = intmath.Atoi(sp)
	}
	s.points = append(s.points, p)
}

func buildScanners(in []string) []scanner {
	ret := []scanner{}

	s := scanner{}
	for i, line := range in {
		if strings.Contains(line, "scanner") {
			if i != 0 {
				ret = append(ret, s)
			}
			s = scanner{id: i}
			continue
		}
		if len(line) > 0 {
			s.addPoints(line)
		}
	}
	ret = append(ret, s)

	return ret
}

func part1(in []string) {
	s := buildScanners(in)
	fmt.Println(s)
}

func part2(in []string) {
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

func input() []string {
	return fileinput.ReadLines("input.txt")
}
