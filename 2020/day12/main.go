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
	inst []string
	pos  intmath.Point
	dir  int
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
			switch (b.dir+4000) % 4 {
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
		fmt.Printf("%v: %v\n", i, b)
	}
}

func part1(b *boat) {
	b.run()
	fmt.Printf("The boat is %v away from start.\n", b.pos.ManhattanDistanceTo(intmath.Point{}))
}

func part2(b *boat) {
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
		inst: fileinput.ReadLines("input.txt"),
		dir:  E,
	}
}
