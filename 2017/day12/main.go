package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func input() (ret [][]string) {
	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, strings.Fields(line))
	}
	return
}

type village struct {
	id    string
	pairs []*village
}

func (v *village) pairWith(ov *village) {
	v.pairs = append(v.pairs, ov)
	ov.pairs = append(ov.pairs, v)
}

func getOrBuildVillage(id string, m map[string]*village) *village {
	v := m[id]
	if v == nil {
		v = &village{id: id}
		m[id] = v
	}
	return v
}

func makeVillages() map[string]*village {
	ret := map[string]*village{}
	for _, line := range input() {
		v := getOrBuildVillage(line[0], ret)
		for _, oid := range line[2:] {
			oid = strings.ReplaceAll(oid, ",", "")
			ov := getOrBuildVillage(oid, ret)
			v.pairWith(ov)
		}
	}
	return ret
}

func sizeOfGroup(v *village, visited map[*village]bool) int {
	var ret int
	if !visited[v] {
		visited[v] = true
		ret++
	}
	for _, ov := range v.pairs {
		if !visited[ov] {
			ret += sizeOfGroup(ov, visited)
		}
	}
	return ret
}

func part1() {
	v := makeVillages()
	fmt.Println("The size of the group with village 0 is:", sizeOfGroup(v["0"], map[*village]bool{}))
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
