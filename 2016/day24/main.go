package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

const (
	keyF = "%d->%d"
)

type maze struct {
	m   [][]rune
	poi map[string]intmath.Point
}

func newMaze() *maze {
	m := maze{}
	for _, line := range fileinput.ReadLines("input.txt") {
		m.m = append(m.m, []rune(line))
	}
	m.markPOI()
	return &m
}

func (m *maze) markPOI() {
	m.poi = map[string]intmath.Point{}
	for y, row := range m.m {
		for x, cell := range row {
			if cell >= '0' && cell <= '7' {
				m.poi[string(cell)] = intmath.Point{X: x, Y: y}
			}
		}
	}
}

type walkState struct {
	p     intmath.Point
	steps int
}

func (m *maze) print() {
	for _, row := range m.m {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}

func (m *maze) minDistance(from, to intmath.Point) int {
	state := walkState{p: from}
	states := []walkState{state}
	for len(states) > 0 {
		state = states[0]
		states = states[1:]
		if m.m[state.p.Y][state.p.X] == '#' {
			continue
		}
		m.m[state.p.Y][state.p.X] = '#'
		for _, n := range state.p.Neighbors() {
			if n == to {
				return state.steps + 1
			}
			states = append(states, walkState{p: n, steps: state.steps + 1})
		}
	}
	return -1
}

func minDistances() map[string]int {
	ret := map[string]int{}
	for i := 0; i <= 7; i++ {
		for j := 0; j <= 7; j++ {
			if i != j {
				if _, ok := ret[fmt.Sprintf(keyF, i, j)]; !ok {
					m := newMaze()
					from := m.poi[fmt.Sprintf("%d", i)]
					to := m.poi[fmt.Sprintf("%d", j)]
					d := m.minDistance(from, to)
					ret[fmt.Sprintf(keyF, i, j)] = d
					ret[fmt.Sprintf(keyF, j, i)] = d
				}
			}
		}
	}
	return ret
}

type path struct {
	p []int
	d int
}

func (p path) String() string {
	return fmt.Sprintf("Path: %v, Distance: %d", p.p, p.d)
}

func part1() {
	mins := minDistances()
	var paths []path
	for perm := range intmath.Permutations([]int{1, 2, 3, 4, 5, 6, 7}) {
		perm = append([]int{0}, perm...)
		p := path{p: perm}
		for i := 0; i < len(perm)-1; i++ {
			p.d += mins[fmt.Sprintf(keyF, perm[i], perm[i+1])]
		}
		paths = append(paths, p)
	}
	sort.Slice(paths, func(i, j int) bool { return paths[i].d < paths[j].d })
	fmt.Println("The minimum path is:", paths[0])
}

func part2() {
	mins := minDistances()
	var paths []path
	for perm := range intmath.Permutations([]int{1, 2, 3, 4, 5, 6, 7}) {
		perm = append([]int{0}, perm...)
		perm = append(perm, 0)
		p := path{p: perm}
		for i := 0; i < len(perm)-1; i++ {
			p.d += mins[fmt.Sprintf(keyF, perm[i], perm[i+1])]
		}
		paths = append(paths, p)
	}
	sort.Slice(paths, func(i, j int) bool { return paths[i].d < paths[j].d })
	fmt.Println("The minimum path returning to 0 is:", paths[0])
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
