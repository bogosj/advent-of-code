package donutmaze

import (
	"github.com/bogosj/advent-of-code/2019/fileinput"
	"github.com/bogosj/advent-of-code/2019/intmath"
)

const (
	space = '.'
	wall  = '#'
)

// Maze represents a donut maze on Pluto.
type Maze struct {
	m       [][]rune
	warps   []intmath.Point
	visited map[intmath.Point]bool
}

// New creates a new maze from a provided text file.
func New(p string) *Maze {
	m := Maze{}
	m.m = input(p)
	m.findWarpPoints()
	m.visited = map[intmath.Point]bool{}
	return &m
}

func (m *Maze) String() (ret string) {
	for _, row := range m.m {
		for _, c := range row {
			ret += string(c)
		}
		ret += "\n"
	}
	return
}

type pointDist struct {
	p    intmath.Point
	dist int
}

// ShortestPath performs a BFS across the maze using warp points and returns the shortest path in steps.
func (m *Maze) ShortestPath() int {
	points := []pointDist{pointDist{p: m.startPoint(), dist: 0}}
	for len(points) > 0 {
		point := points[0]
		points = points[1:]
		if _, ok := m.visited[point.p]; ok {
			continue
		}
		m.visited[point.p] = true
		if point.p == m.endPoint() {
			return point.dist
		}
		if m.isWarpPoint(point.p) {
			points = append(points, pointDist{p: m.warpPointOtherEnd(point.p), dist: point.dist + 1})
		}
		for _, n := range point.p.Neighbors() {
			if m.m[n.Y][n.X] == space {
				points = append(points, pointDist{p: n, dist: point.dist + 1})
			}
		}
	}
	return -1
}

func (m *Maze) isLetter(p intmath.Point) bool {
	return m.m[p.Y][p.X] >= 'A' && m.m[p.Y][p.X] <= 'Z'
}

func (m *Maze) isWarpPoint(p intmath.Point) bool {
	if p == m.startPoint() || p == m.endPoint() {
		return false
	}
	for _, w := range m.warps {
		if w == p {
			return true
		}
	}
	return false
}

func (m *Maze) warpPointsByName(n string) (ret []intmath.Point) {
	for _, w := range m.warps {
		if m.warpPointName(w) == n {
			ret = append(ret, w)
		}
	}
	return
}

func (m *Maze) warpPointOtherEnd(p intmath.Point) intmath.Point {
	n := m.warpPointName(p)
	for _, op := range m.warpPointsByName(n) {
		if op != p {
			return op
		}
	}
	return intmath.Point{}
}

func (m *Maze) startPoint() (ret intmath.Point) {
	w := m.warpPointsByName("AA")
	return w[0]
}

func (m *Maze) endPoint() (ret intmath.Point) {
	w := m.warpPointsByName("ZZ")
	return w[0]
}

func (m *Maze) warpPointName(p intmath.Point) string {
	var ret []rune
	for _, p1 := range p.Neighbors() {
		if m.isLetter(p1) {
			ret = append(ret, m.m[p1.Y][p1.X])
			for _, p2 := range p1.Neighbors() {
				if m.isLetter(p2) {
					ret = append(ret, m.m[p2.Y][p2.X])
				}
			}
		}
	}
	if ret[0] < ret[1] {
		return string(ret[0]) + string(ret[1])
	}
	return string(ret[1]) + string(ret[0])
}

func (m *Maze) findWarpPoints() {
	for y := range m.m {
		for x := range m.m[y] {
			if m.m[y][x] == space {
				p := intmath.Point{X: x, Y: y}
				for _, n := range p.Neighbors() {
					if m.isLetter(n) {
						m.warps = append(m.warps, p)
					}
				}
			}
		}
	}
}

func input(p string) (ret [][]rune) {
	lines := fileinput.ReadLinesRaw(p)
	for _, line := range lines {
		row := []rune{}
		for _, r := range line {
			row = append(row, r)
		}
		ret = append(ret, row)
	}
	return
}
