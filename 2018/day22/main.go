package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/intmath"
)

type point = intmath.Point
type regionType = int

const (
	rocky = iota
	wet
	narrow
)

type cave struct {
	geoIndexes map[point]int
	depth      int
	target     point
}

func newCave() *cave {
	return &cave{
		depth:      3558,
		target:     point{X: 15, Y: 740},
		geoIndexes: map[point]int{},
	}
}

func (c *cave) erosionLevel(p point) int {
	return (c.geoIndex(p) + c.depth) % 20183
}

func (c *cave) regionType(p point) regionType {
	return c.erosionLevel(p) % 3
}

func (c *cave) geoIndex(p point) (ret int) {
	switch v, ok := c.geoIndexes[p]; {
	case ok:
		ret = v
	case p.ManhattanDistanceTo(point{}) == 0:
	case p.ManhattanDistanceTo(c.target) == 0:
	case p.Y == 0:
		ret = p.X * 16807
	case p.X == 0:
		ret = p.Y * 48271
	default:
		ret = c.erosionLevel(p.Neighbor('L')) * c.erosionLevel(p.Neighbor('U'))
	}
	c.geoIndexes[p] = ret
	return
}

func (c *cave) riskLevel() (risk int) {
	for y := 0; y <= c.target.Y; y++ {
		for x := 0; x <= c.target.X; x++ {
			risk += c.regionType(point{X: x, Y: y})
		}
	}
	return
}

func part1() {
	c := newCave()
	fmt.Println("The risk level of this quadrant is:", c.riskLevel())
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
