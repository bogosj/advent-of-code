package beam

import (
	"fmt"

	"github.com/bogosj/advent-of-code/2019/computer"
)

// Beam represents a tractor beam.
type Beam struct {
	c *computer.Computer
	m [][]rune
}

// New creates a new tractor beam with the provided computer.
func New(c *computer.Computer) *Beam {
	b := Beam{}
	b.c = c
	b.c.OneOutput = true
	return &b
}

func (b *Beam) scanLoc(x, y int) int {
	b.c.Reset()
	in := make(chan int)
	out := b.c.Compute(in)
	in <- x
	in <- y
	return <-out
}

// Scan runs the computer program to scan an area.
func (b *Beam) Scan() (ret int) {
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			output := b.scanLoc(x, y)
			if output == 1 {
				ret++
			}
			fmt.Print(output)
		}
		fmt.Println()
	}
	return
}

// ScanFor10x10 finds the closest point of a 10x10 contiguous hull.
func (b *Beam) ScanFor10x10() (x, y int) {
	x = 0
	y = 15
	for {
		output := b.scanLoc(x, y)
		if output == 0 {
			x++
			continue
		}
		if b.scanLoc(x, y-99) == 1 && b.scanLoc(x+99, y-99) == 1 {
			y -= 99
			return
		}
		y++
	}
}
