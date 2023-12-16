package network

import (
	"time"

	"github.com/bogosj/advent-of-code/2019/computer"
)

// Network represents a network on the ship.
type Network struct {
	ins       []chan<- int
	outs      []<-chan int
	computers []*computer.Computer
	natVal    *instruction
}

// New creates a new network of size s, with a program at path p.
func New(s int, p string) *Network {
	n := Network{}
	for i := 0; i < s; i++ {
		c := computer.New(p)
		in := make(chan int, 500)
		out := c.Compute(in)
		n.ins = append(n.ins, in)
		n.outs = append(n.outs, out)
		n.computers = append(n.computers, c)
		in <- i
	}
	return &n
}

type instruction struct {
	a, x, y int
}

// Run allows the network to run, and returns the first value sent to address 255.
func (n *Network) Run(returnFirst bool) int {
	var buffer []instruction
	natYs := map[int]bool{}
	for i := 0; i < len(n.ins); i++ {
		n.ins[i] <- -1
	}
	for {
		time.Sleep(time.Millisecond * 5)
		for i := 0; i < len(n.outs); i++ {
			out := n.outs[i]
			select {
			case a := <-out:
				x := <-out
				y := <-out
				buffer = append(buffer, instruction{a: a, x: x, y: y})
			default:
			}
		}
		for _, inst := range buffer {
			if inst.a == 255 {
				if returnFirst && n.natVal == nil {
					return inst.y
				}
				n.natVal = &inst
			} else {
				n.ins[inst.a] <- inst.x
				n.ins[inst.a] <- inst.y
			}
		}
		if len(buffer) == 0 {
			i := n.natVal
			n.ins[0] <- i.x
			n.ins[0] <- i.y
			if _, ok := natYs[i.y]; ok {
				return i.y
			}
			natYs[i.y] = true
		}
		buffer = nil
	}
}
