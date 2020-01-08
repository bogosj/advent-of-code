package main

import (
	"fmt"
	"time"

	"gonum.org/v1/gonum/stat/combin"
)

type generators struct {
	elevator int
	things   []int
	moves    int
}

func newGen() generators {
	g := generators{}
	g.elevator = 1
	// the first five indexes are generators, the second five are chips.
	g.things = []int{1, 1, 1, 3, 3, 1, 2, 2, 3, 3}
	return g
}

func (g generators) copyState() generators {
	ng := generators{
		elevator: g.elevator,
		moves:    g.moves + 1,
	}
	for _, v := range g.things {
		ng.things = append(ng.things, v)
	}
	return ng
}

func (g generators) final() bool {
	for _, v := range g.things {
		if v != 4 {
			return false
		}
	}
	return true
}

func (g generators) moveIdxs(idxs []int, combos [][]int) (ret []generators) {
	for _, c := range combos {
		if g.elevator < 4 {
			ng := g.copyState()
			ng.elevator++
			for _, idx := range c {
				ng.things[idxs[idx]]++
			}
			ret = append(ret, ng)
		}
		if g.elevator > 1 {
			ng := g.copyState()
			ng.elevator--
			for _, idx := range c {
				ng.things[idxs[idx]]--
			}
			ret = append(ret, ng)
		}
	}
	return
}

func (g generators) nextStates() (ret []generators) {
	var idxs []int
	for idx, v := range g.things {
		if v == g.elevator {
			idxs = append(idxs, idx)
		}
	}
	if len(idxs) > 0 {
		ngs := g.moveIdxs(idxs, combin.Combinations(len(idxs), 1))
		ret = append(ret, ngs...)
	}
	if len(idxs) > 1 {
		ngs := g.moveIdxs(idxs, combin.Combinations(len(idxs), 2))
		ret = append(ret, ngs...)
	}
	return
}

func (g generators) safeState() bool {
	for idx, floor := range g.things[len(g.things)/2:] {
		// if the chip is on the same floor as its generator, it is safe
		if floor == g.things[idx] {
			continue
		}
		// if there is a generator of a different type on the same floor, it's not safe
		for genIdx, genFloor := range g.things[:len(g.things)/2] {
			if genIdx == idx {
				continue
			}
			if genFloor == floor {
				return false
			}
		}
	}
	return true
}

func minMoves(start generators) int {
	states := []generators{start}
	seenStates := map[string]bool{}
	for len(states) > 0 {
		state := states[0]
		states = states[1:]
		stateStr := fmt.Sprintf("%v|%v", state.elevator, state.things)
		if seenStates[stateStr] {
			continue
		}
		seenStates[stateStr] = true
		if !state.safeState() {
			continue
		}
		if state.final() {
			return state.moves
		}
		ns := state.nextStates()
		states = append(states, ns...)
	}
	return -1
}

func part1() {
	g := newGen()
	fmt.Println("The minimum number of moves is:", minMoves(g))
}

func part2() {
	g := newGen()
	g.things = []int{1, 1, 1, 3, 3, 1, 1, 1, 2, 2, 3, 3, 1, 1}
	fmt.Println("The minimum number of moves with extra stuff is:", minMoves(g))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
