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

// Scan runs the computer program to scan an area.
func (b *Beam) Scan() (ret int) {
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			b.c.Reset()
			in := make(chan int, 2)
			out := b.c.Compute(in)
			in <- x
			in <- y
			output := <-out
			if output == 1 {
				ret++
			}
			fmt.Print(output)
		}
		fmt.Println()
	}
	return
}
