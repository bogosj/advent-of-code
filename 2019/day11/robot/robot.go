package robot

import "fmt"

import "github.com/bogosj/advent-of-code/2019/computer"

const (
	dirUp = iota
	dirRight
	dirDown
	dirLeft
	white = '#'
	black = '.'
)

type Robot struct {
	c       *computer.Computer
	hull    [][]rune
	x, y, d int
}

func New(c *computer.Computer) *Robot {
	r := Robot{}
	r.c = c
	return &r
}

func (r *Robot) readCurrentPaint() int {
	if r.hull[r.x][r.y] == white {
		return 1
	}
	return 0
}

func (r *Robot) PrintHull() (ret int) {
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

func (r *Robot) turn(d int) {
	switch d {
	case 0:
		r.d--
	case 1:
		r.d++
	}
	r.d = (r.d + 4) % 4
}

func (r *Robot) move() {
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

func (r *Robot) Paint(start int) {
	r.hull = make([][]rune, 300)
	for i := range r.hull {
		r.hull[i] = make([]rune, 300)
	}
	r.x = 70
	r.y = 70

	out := r.c.Compute(start)
	for !r.c.Halted {
		if out == 1 {
			r.hull[r.x][r.y] = white
		} else {
			r.hull[r.x][r.y] = black
		}
		out = r.c.Compute(0)
		r.turn(out)
		r.move()
		out = r.c.Compute(r.readCurrentPaint())
	}
}
