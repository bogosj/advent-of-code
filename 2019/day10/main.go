package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func absgcd(a, b int) int {
	a = abs(a)
	b = abs(b)
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

type point struct {
	x, y int
}

type starmap struct {
	m [][]rune
}

func (s *starmap) width() int {
	return len(s.m)
}

func (s *starmap) height() int {
	return len(s.m[0])
}

func (s *starmap) lineOfSight(from, to point) bool {
	xd, yd := delta(from, to)
	x := from.x
	y := from.y
	for {
		x += xd
		y += yd
		if x == to.x && y == to.y {
			return true
		}
		if s.m[x][y] == '#' {
			return false
		}
	}
}

func countLineOfSight(m starmap, to point) (ret int) {
	for x := 0; x < m.width(); x++ {
		for y := 0; y < m.height(); y++ {
			if x == to.x && y == to.y {
				continue
			}
			if m.m[x][y] == '#' {
				if m.lineOfSight(point{x, y}, point{to.x, to.y}) {
					ret++
				}
			}
		}
	}
	return
}

func delta(from, to point) (xd, yd int) {
	xd = to.x - from.x
	yd = to.y - from.y
	if xd == 0 {
		yd /= abs(yd)
		return
	}
	if yd == 0 {
		xd /= abs(xd)
		return
	}
	gcd := absgcd(xd, yd)
	xd /= gcd
	yd /= gcd
	return
}

func part1() {
	m := starmap{input("input.txt")}
	max := 0
	var station point
	for x := 0; x < m.width(); x++ {
		for y := 0; y < m.height(); y++ {
			if m.m[x][y] == '#' {
				los := countLineOfSight(m, point{x, y})
				if los > max {
					max = los
					station = point{x, y}
				}
			}
		}
	}
	fmt.Println("max asteroids:", max, "From station at", station)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("part 1 complete in:", time.Since(start))
}

func input(n string) (ret [][]rune) {
	lines := strings.Split(rawinput(n), "\n")
	for _, line := range lines {
		ret = append(ret, []rune(line))
	}
	return ret
}

func rawinput(n string) string {
	data, _ := ioutil.ReadFile(n)
	return string(data)
}
