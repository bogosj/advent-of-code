package camera

import "github.com/bogosj/advent-of-code/2019/computer"

import "github.com/bogosj/advent-of-code/2019/intmath"

const (
	scaffold  = '#'
	openSpace = '.'
)

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
