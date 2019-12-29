package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

// Point defines a point in cartesian space.
type Point struct {
	X, Y int
}

func (p *Point) dist() int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
}

// Wire defines the path of a wire.
type Wire struct {
	w map[Point]int
	p Point
	l int
}

// New creates a new Wire
func New() Wire {
	w := Wire{}
	w.w = map[Point]int{}
	return w
}

func (w *Wire) move(m string) {
	dir := string(m[0])
	dist, _ := strconv.Atoi(m[1:])
	x := w.p.X
	y := w.p.Y
	for i := 0; i < dist; i++ {
		switch dir {
		case "R":
			y++
		case "L":
			y--
		case "U":
			x++
		case "D":
			x--
		}
		w.p = Point{x, y}
		w.l++
		if w.w[w.p] == 0 {
			w.w[w.p] = w.l
		}
	}
}

func (w *Wire) intersects(w2 Wire) []Point {
	var ret []Point
	for k := range w.w {
		if _, ok := w2.w[k]; ok {
			ret = append(ret, k)
		}
	}
	return ret
}

func (w *Wire) distanceOnWire(p Point) int {
	return w.w[p]
}

func closestPointManhattan(points []Point) {
	var dists []int
	for _, p := range points {
		dists = append(dists, p.dist())
	}
	sort.Ints(dists)
	fmt.Printf("closest point: %v\n", dists[0])
}

func closestPointOnWires(points []Point, w1, w2 Wire) {
	minDist := math.MaxInt64
	for _, p := range points {
		dist := w1.distanceOnWire(p) + w2.distanceOnWire(p)
		if dist < minDist {
			minDist = dist
		}
	}
	fmt.Printf("closest point on wire: %v\n", minDist)
}

func main() {
	i := input()
	w1 := New()
	for _, move := range i[0] {
		w1.move(move)
	}
	w2 := New()
	for _, move := range i[1] {
		w2.move(move)
	}
	points := w1.intersects(w2)
	start := time.Now()
	closestPointManhattan(points)
	fmt.Println("part 1 done in:", time.Since(start))
	start = time.Now()
	closestPointOnWires(points, w1, w2)
	fmt.Println("part 2 done in:", time.Since(start))
}

func input() [][]string {
	lines := fileinput.ReadLines("input.txt")
	w1 := strings.Split(lines[0], `,`)
	w2 := strings.Split(lines[1], `,`)
	return [][]string{w1, w2}
}
