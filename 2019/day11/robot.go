package main

import "fmt"

const (
	dirUp = iota
	dirRight
	dirDown
	dirLeft
	white = '#'
	black = '.'
)

type robot struct {
	c       *computer
	hull    [][]rune
	x, y, d int
}

func (r *robot) readCurrentPaint() int {
	if r.hull[r.x][r.y] == white {
		return 1
	}
	return 0
}

func (r *robot) printHull() (ret int) {
	for i := range r.hull {
		for j := range r.hull[i] {
			if r.hull[i][j] == white || r.hull[i][j] == black {
				ret++
			}
			if r.hull[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(string(r.hull[i][j]))
			}
		}
		fmt.Println()
	}
	return
}

func (r *robot) turn(d int) {
	switch d {
	case 0:
		r.d--
	case 1:
		r.d++
	}
	r.d = (r.d + 4) % 4
}

func (r *robot) move() {
	switch r.d {
	case dirUp:
		r.x++
	case dirDown:
		r.x--
	case dirLeft:
		r.y--
	case dirRight:
		r.y++
	}
}

func (r *robot) paint(start int) {
	r.hull = make([][]rune, 300)
	for i := range r.hull {
		r.hull[i] = make([]rune, 300)
	}
	r.x = 70
	r.y = 70

	var out int
	var err error
	out, err = r.c.compute(start)
	for err == nil {
		if out == 1 {
			r.hull[r.x][r.y] = white
		} else {
			r.hull[r.x][r.y] = black
		}
		out, err = r.c.compute(0)
		r.turn(out)
		r.move()
		out, err = r.c.compute(r.readCurrentPaint())
	}
}
