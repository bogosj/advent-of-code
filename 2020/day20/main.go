package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type tile struct {
	id   int
	data []string
}

func reverse(s string) (ret string) {
	for _, c := range s {
		ret = string(c) + ret
	}
	return
}

func (t *tile) sides() (ret []string) {
	ret = append(ret, t.data[0], reverse(t.data[0]))
	ret = append(ret, t.data[len(t.data)-1], reverse(t.data[len(t.data)-1]))
	s1, s2, l := "", "", len(t.data[0])-1
	for y := 0; y < len(t.data); y++ {
		s1 += string(t.data[y][0])
		s2 += string(t.data[y][l])
	}
	ret = append(ret, s1, reverse(s1))
	ret = append(ret, s2, reverse(s2))
	return
}

func part1(in []tile) {
	stringToTile := map[string][]tile{}
	for _, t := range in {
		for _, id := range t.sides() {
			stringToTile[id] = append(stringToTile[id], t)
		}
	}
	edgeTiles := map[int]int{}
	for _, v := range stringToTile {
		if len(v) == 1 {
			edgeTiles[v[0].id]++
		}
	}
	product := 1
	for k, v := range edgeTiles {
		if v == 4 {
			product *= k
		}
	}
	fmt.Printf("Product of corner tile IDs is %v\n", product)
}

func part2(in []tile) {
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

func input() []tile {
	ret := []tile{}

	var t tile
	for _, line := range fileinput.ReadLines("input.txt") {
		f := strings.Fields(line)
		if line == "" {
			ret = append(ret, t)
		} else if f[0] == "Tile" {
			t = tile{}
			sid := strings.ReplaceAll(f[1], ":", "")
			t.id = intmath.Atoi(sid)
		} else {
			t.data = append(t.data, line)
		}
	}

	return ret
}
