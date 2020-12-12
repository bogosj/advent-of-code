package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

const (
	E = iota
	S
	W
	N
)

type boat struct {
	inst     []string
	pos      intmath.Point
	dir      int
	waypoint intmath.Point
}

func (b *boat) String() string {
	return fmt.Sprintf("pos: %v, dir: %v", b.pos, b.dir)
}

func (b *boat) run() {
	for _, i := range b.inst {
		r := i[0]
		n := intmath.Atoi(i[1:])
		switch r {
		case 'N':
			b.pos.Y += n
		case 'S':
			b.pos.Y -= n
		case 'E':
			b.pos.X += n
		case 'W':
			b.pos.X -= n
		case 'L':
			b.dir -= n / 90
		case 'R':
			b.dir += n / 90
		case 'F':
			switch (b.dir + 4000) % 4 {
			case N:
				b.pos.Y += n
			case S:
				b.pos.Y -= n
			case E:
				b.pos.X += n
			case W:
				b.pos.X -= n
			}
		}
	}
}

func (b *boat) rotate(n int, d rune) {
	if d == 'L' {
		for i := 0; i < n; i++ {
			b.waypoint.Y *= -1
			b.waypoint = intmath.Point{X: b.waypoint.Y, Y: b.waypoint.X}
		}
	} else {
		for i := 0; i < n; i++ {
			b.waypoint.X *= -1
			b.waypoint = intmath.Point{X: b.waypoint.Y, Y: b.waypoint.X}
		}
	}
}

func (b *boat) run2() {
	for _, i := range b.inst {
		r := i[0]
		n := intmath.Atoi(i[1:])
		switch r {
		case 'N':
			b.waypoint.Y += n
		case 'S':
			b.waypoint.Y -= n
		case 'E':
			b.waypoint.X += n
		case 'W':
			b.waypoint.X -= n
		case 'L', 'R':
			b.rotate(n/90, rune(r))
		case 'F':
			b.pos.Y += n * b.waypoint.Y
			b.pos.X += n * b.waypoint.X
		}
	}
}

func part1(b *boat) {
	b.run()
	fmt.Printf("The boat is %v away from start.\n", b.pos.ManhattanDistanceTo(intmath.Point{}))
}

func part2(b *boat) {
	b.run2()
	fmt.Printf("The boat is %v away from start.\n", b.pos.ManhattanDistanceTo(intmath.Point{}))
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

func input() *boat {
	return &boat{
		inst:     fileinput.ReadLines("input.txt"),
		dir:      E,
		waypoint: intmath.Point{X: 10, Y: 1},
	}
}
