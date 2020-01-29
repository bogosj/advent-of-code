package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type nanobot struct {
	x, y, z, r int
}

func (n *nanobot) String() string {
	return fmt.Sprintf("<%d %d %d> | %d", n.x, n.y, n.z, n.r)
}

func (n *nanobot) canSee(on *nanobot) bool {
	dist := intmath.Abs(n.x-on.x) + intmath.Abs(n.y-on.y) + intmath.Abs(n.z-on.z)
	return n.r >= dist
}

func input() (ret []*nanobot) {
	for _, line := range fileinput.ReadLines("input.txt") {
		b := &nanobot{}
		f := strings.Split(line, " ")
		b.r = intmath.Atoi(strings.Split(f[1], "=")[1])
		p := strings.Split(f[0], "=")[1]
		p = strings.ReplaceAll(p, "<", "")
		p = strings.ReplaceAll(p, ">", "")
		pos := strings.Split(p, ",")
		b.x, b.y, b.z = intmath.Atoi(pos[0]), intmath.Atoi(pos[1]), intmath.Atoi(pos[2])
		ret = append(ret, b)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].r > ret[j].r
	})
	return
}

func part1() {
	bots := input()
	var c int
	for _, bot := range bots {
		if bots[0].canSee(bot) {
			c++
		}
	}
	fmt.Printf("There are %d bots in range of the bot with the biggest range.\n", c)
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
