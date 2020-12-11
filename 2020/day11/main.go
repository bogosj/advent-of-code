package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/intmath"

	"github.com/bogosj/advent-of-code/fileinput"
)

type boat struct {
	rows [][]rune
}

func (b *boat) occupiedSeats() (ret int) {
	for _, row := range b.rows {
		for _, c := range row {
			if c == '#' {
				ret++
			}
		}
	}
	return
}

func (b *boat) seatOccupied(p intmath.Point) bool {
	if p.Y < 0 || p.Y >= len(b.rows) {
		return false
	}
	if p.X < 0 || p.X >= len(b.rows[0]) {
		return false
	}
	return b.rows[p.Y][p.X] == '#'
}

func (b *boat) next() (changed bool) {
	newState := [][]rune{}
	for y, row := range b.rows {
		nr := []rune{}
		for x, seat := range row {
			p := intmath.Point{X: x, Y: y}
			c := 0
			for _, n := range p.AllNeighbors() {
				if b.seatOccupied(n) {
					c++
				}
			}
			if seat == 'L' && c == 0 {
				nr = append(nr, '#')
				changed = true
			} else if seat == '#' && c >= 4 {
				nr = append(nr, 'L')
				changed = true
			} else {
				nr = append(nr, seat)
			}
		}
		newState = append(newState, nr)
	}
	b.rows = newState
	return
}

func part1(b boat) {
	i := 0
	for {
		i++
		if !b.next() {
			fmt.Printf("It took %v iterations to find the layout of %v seats.\n", i, b.occupiedSeats())
			return
		}
	}
}

func part2(b boat) {
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() boat {
	b := boat{}
	for _, line := range fileinput.ReadLines("input.txt") {
		b.rows = append(b.rows, []rune(line))
	}
	return b
}
