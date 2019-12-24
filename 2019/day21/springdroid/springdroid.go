package springdroid

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/computer"
	"github.com/bogosj/advent-of-code/2019/fileinput"
)

// Droid represents a SpringDroid
type Droid struct {
	c *computer.Computer
}

// New creates a new SpringDroid with a computer defined by the program at the provided path.
func New(p string) *Droid {
	d := Droid{}
	d.c = computer.New(p)
	return &d
}

// RunProgram runs the program from disk at the provided path.
func (d *Droid) RunProgram(p string) {
	prog := fileinput.ReadLines(p)
	in := make(chan int)
	out := d.c.Compute(in)
READ:
	for {
		select {
		case o := <-out:
			fmt.Print(string(o))
		case <-time.After(time.Millisecond * 30):
			break READ
		}
	}
	for _, line := range prog {
		for _, r := range line {
			in <- int(r)
		}
		in <- 10
	}
	for o := range out {
		if o > 255 {
			fmt.Printf("Output: %d\n", o)
		} else {
			fmt.Print(string(o))
		}
	}
}
