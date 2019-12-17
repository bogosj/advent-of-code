package camera

import (
	"fmt"

	"github.com/bogosj/advent-of-code/2019/computer"
	"github.com/bogosj/advent-of-code/2019/intmath"
)

const (
	scaffold  = '#'
	openSpace = '.'
	moveMain  = "A,B,B,C,C,A,A,B,B,C"
	moveA     = "L,12,R,4,R,4"
	moveB     = "R,12,R,4,L,12"
	moveC     = "R,12,R,4,L,6,L,8,L,8"
)

func moveInst(in string) (ret []int) {
	for _, c := range in {
		ret = append(ret, int(c))
	}
	ret = append(ret, 10)
	return
}

// Camera represents an ASCII camera.
type Camera struct {
	c *computer.Computer
	i [][]rune
}

// New creates a new camera with a provided computer.
func New(c *computer.Computer) (ret Camera) {
	ret.c = c
	return
}

func (c *Camera) readOutput() {
	var prev int
	for {
		out := c.c.Compute()
		fmt.Print(string(out))
		if (prev == ':' || prev == '?') && out == '\n' {
			return
		}
		prev = out
	}
}

// Notify wakes up all other robots and cleans them, returning the amount of dust cleaned.
func (c *Camera) Notify() int {
	c.c.Hack(0, 2)

	for _, s := range []string{moveMain, moveA, moveB, moveC, "n\n"} {
		c.readOutput()
		mi := moveInst(s)
		fmt.Print(string(c.c.Compute(mi...)))
	}
	var prev int
	for !c.c.Halted {
		out := c.c.Compute()
		if out == 0 {
			return prev
		}
		if out < 128 {
			fmt.Print(string(out))
		}
		prev = out
	}
	return 0
}

// CaptureImage returns a string representation of the image captured.
func (c *Camera) CaptureImage() {
	row := []rune{}
	for !c.c.Halted {
		out := c.c.Compute()
		row = append(row, rune(out))

		if out == 10 {
			if row[0] != 10 {
				c.i = append(c.i, row)
				row = []rune{}
			}
		}
	}
	return
}

// Image returns the state of the camera as a string.
func (c *Camera) Image() (ret string) {
	for _, row := range c.i {
		for _, cell := range row {
			ret += string(cell)
		}
	}
	return
}

func (c *Camera) allNeighborsScaffold(p intmath.Point) bool {
	if c.i[p.Y][p.X-1] != scaffold {
		return false
	}
	if c.i[p.Y][p.X+1] != scaffold {
		return false
	}
	if c.i[p.Y-1][p.X] != scaffold {
		return false
	}
	if c.i[p.Y+1][p.X] != scaffold {
		return false
	}
	return true
}

// Intersections returns the points in the image where the scaffold intersects with itself.
func (c *Camera) Intersections() (ret []intmath.Point) {
	for y := 1; y < len(c.i)-1; y++ {
		for x := 1; x < len(c.i[0])-1; x++ {
			p := intmath.Point{X: x, Y: y}
			if c.i[y][x] == scaffold {
				if c.allNeighborsScaffold(p) {
					ret = append(ret, p)
				}
			}
		}
	}
	return
}
