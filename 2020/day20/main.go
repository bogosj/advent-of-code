package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type tile struct {
	id   int
	data []string
	sIDs []int
}

func reverse(s string) (ret string) {
	for _, c := range s {
		ret = string(c) + ret
	}
	return
}

func idsFromString(s string) (ret []int) {
	s = strings.ReplaceAll(s, "#", "1")
	s = strings.ReplaceAll(s, ".", "0")
	i, _ := strconv.ParseInt(s, 2, 0)
	ret = append(ret, int(i))
	i, _ = strconv.ParseInt(reverse(s), 2, 0)
	ret = append(ret, int(i))
	return
}

func (t *tile) sideIDs() []int {
	if t.sIDs != nil {
		return t.sIDs
	}
	t.sIDs = append(t.sIDs, idsFromString(t.data[0])...)
	t.sIDs = append(t.sIDs, idsFromString(t.data[len(t.data)-1])...)
	s1, s2, l := "", "", len(t.data[0])-1
	for y := 0; y < len(t.data); y++ {
		s1 += string(t.data[y][0])
		s2 += string(t.data[y][l])
	}
	t.sIDs = append(t.sIDs, idsFromString(s1)...)
	t.sIDs = append(t.sIDs, idsFromString(s2)...)
	return t.sIDs
}

func part1(in []tile) {
	idToTile := map[int][]tile{}
	for _, t := range in {
		for _, id := range t.sideIDs() {
			idToTile[id] = append(idToTile[id], t)
		}
	}
	edgeTiles := map[int]int{}
	for _, v := range idToTile {
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
