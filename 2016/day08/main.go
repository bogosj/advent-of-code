package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

const (
	on  = '#'
	off = '.'
)

type display struct {
	d [][]rune
}

func newDisplay() *display {
	d := display{}
	for i := 0; i < 6; i++ {
		var row []rune
		for j := 0; j < 50; j++ {
			row = append(row, off)
		}
		d.d = append(d.d, row)
	}
	return &d
}

func (d *display) rect(amt string) {
	f := strings.FieldsFunc(amt, func(r rune) bool {
		return r == 'x'
	})
	for y := 0; y < intmath.Atoi(f[1]); y++ {
		for x := 0; x < intmath.Atoi(f[0]); x++ {
			d.d[y][x] = on
		}
	}
}

func (d *display) column(idx, amt string) {
	f := strings.FieldsFunc(idx, func(r rune) bool {
		return r == '='
	})
	c := intmath.Atoi(f[1])
	a := intmath.Atoi(amt)
	for n := 0; n < a; n++ {
		nv := make([]rune, 6)
		for r := 0; r < 6; r++ {
			nv[(r+1)%6] = d.d[r][c]
		}
		for r := 0; r < 6; r++ {
			d.d[r][c] = nv[r]
		}
	}
}

func (d *display) row(idx, amt string) {
	f := strings.FieldsFunc(idx, func(r rune) bool {
		return r == '='
	})
	r := intmath.Atoi(f[1])
	a := intmath.Atoi(amt)
	for n := 0; n < a; n++ {
		nv := make([]rune, 50)
		for c := 0; c < 50; c++ {
			nv[(c+1)%50] = d.d[r][c]
		}
		for c := 0; c < 50; c++ {
			d.d[r][c] = nv[c]
		}
	}
}

func (d *display) execute(s string) {
	f := strings.Fields(s)
	switch f[0] {
	case "rect":
		d.rect(f[1])
	case "rotate":
		switch f[1] {
		case "row":
			d.row(f[2], f[4])
		case "column":
			d.column(f[2], f[4])
		}
	}
}

func (d *display) print() {
	for _, row := range d.d {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}

func (d *display) run(p string) {
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		d.execute(line)
	}
}

func (d *display) litPixels() (ret int) {
	for _, row := range d.d {
		for _, cell := range row {
			if cell == on {
				ret++
			}
		}
	}
	return
}

func part1() {
	d := newDisplay()
	d.run("input.txt")
	fmt.Println("The number of lit pixels is:", d.litPixels())
}

func part2() {
	d := newDisplay()
	d.run("input.txt")
	d.print()
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
