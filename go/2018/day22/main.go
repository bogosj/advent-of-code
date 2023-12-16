package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/intmath"
	pq "github.com/jupp0r/go-priority-queue"
)

type point = intmath.Point
type regionType = int

const (
	rocky = iota
	wet
	narrow
)

const (
	neither = iota
	torch
	gear
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
		return v
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

type state struct {
	pos     point
	minutes int
	equip   int
}

func (s *state) String() string {
	return fmt.Sprintf("%v: %v", s.pos, s.minutes)
}

func (s *state) isBad(c *cave) bool {
	return s.equip == c.regionType(s.pos)
}

func (c *cave) nextStates(curr *state) (ret []*state) {
	for _, n := range curr.pos.Neighbors() {
		if n.X < 0 || n.Y < 0 {
			continue
		}
		ret = append(ret, &state{pos: n, minutes: curr.minutes + 1, equip: curr.equip})
		for i := 1; i <= 2; i++ {
			ns := &state{pos: curr.pos, minutes: curr.minutes + 8, equip: (curr.equip + i) % 3}
			if !ns.isBad(c) {
				ns.pos = n
				ret = append(ret, ns)
			}
		}
	}
	return
}

func (c *cave) navigate() *state {
	states := pq.New()
	states.Insert(&state{equip: torch}, 0)
	seen := map[state]bool{}
	for states.Len() > 0 {
		i, err := states.Pop()
		if err != nil {
			panic(err)
		}
		curr := i.(*state)
		if curr.isBad(c) {
			continue
		}
		cs := state{pos: curr.pos, equip: curr.equip}
		if seen[cs] == true {
			continue
		}
		seen[cs] = true
		if curr.pos.ManhattanDistanceTo(c.target) == 0 {
			if curr.equip == torch {
				return curr
			}
			curr.equip = torch
			curr.minutes += 7
			states.Insert(curr, float64(curr.minutes))
			continue
		}
		for _, s := range c.nextStates(curr) {
			states.Insert(s, float64(s.minutes))
		}
	}
	return &state{}
}

func part1() {
	c := newCave()
	fmt.Println("The risk level of this quadrant is:", c.riskLevel())
}

func part2() {
	c := newCave()
	fmt.Println("The shortest path to the target in minutes is:", c.navigate().minutes)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
