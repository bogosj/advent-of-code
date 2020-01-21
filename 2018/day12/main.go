package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

const (
	start = "####..##.##..##..#..###..#....#.######..###########.#...#.##..####.###.#.###.###..#.####..#.#..##..#"
)

func input() map[string]rune {
	ret := map[string]rune{}
	for _, line := range fileinput.ReadLines("input.txt") {
		f := strings.Fields(line)
		ret[f[0]] = rune(f[2][0])
	}
	return ret
}

func startState() map[int]rune {
	s := map[int]rune{}
	for i, v := range start {
		s[i] = v
	}
	return s
}

func min(s map[int]rune) (ret int) {
	ret = 10
	for i := range s {
		if i < ret {
			ret = i
		}
	}
	return
}

func max(s map[int]rune) (ret int) {
	ret = 0
	for i := range s {
		if i > ret {
			ret = i
		}
	}
	return
}

func nextState(i int, s map[int]rune, t map[string]rune) rune {
	var rs []rune
	for j := i - 2; j <= i+2; j++ {
		if r, ok := s[j]; ok {
			rs = append(rs, r)
		} else {
			rs = append(rs, '.')
		}
	}
	return t[string(rs)]
}

func evolve(s map[int]rune, t map[string]rune) map[int]rune {
	ns := map[int]rune{}
	start, end := min(s)-2, max(s)+2
	for i := start; i <= end; i++ {
		ns[i] = nextState(i, s, t)
	}
	return ns
}

func part1() {
	transitions := input()
	state := startState()
	for i := 0; i < 20; i++ {
		state = evolve(state, transitions)
	}
	var sum int
	for i, v := range state {
		if v == '#' {
			sum += i
		}
	}
	fmt.Println("The sum of all pots holding a plant after 20 generations is:", sum)
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
