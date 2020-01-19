package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type component struct {
	name string
}

func (c component) ports() (a, b string) {
	f := strings.FieldsFunc(c.name, func(r rune) bool { return r == '/' })
	return f[0], f[1]
}

func (c component) hasPort(pins string) bool {
	a, b := c.ports()
	return a == pins || b == pins
}

func (c component) otherPort(pins string) string {
	a, b := c.ports()
	if a == pins {
		return b
	}
	return a
}

func (c component) strength() (ret int) {
	a, b := c.ports()
	return intmath.Atoi(a) + intmath.Atoi(b)
}

func (c component) in(b bridge) bool {
	for _, bc := range b.cs {
		if bc.name == c.name {
			return true
		}
	}
	return false
}

func allComponents() (ret []component) {
	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, component{name: line})
	}
	return
}

type bridge struct {
	cs []component
}

func (b bridge) strength() (ret int) {
	for _, c := range b.cs {
		ret += c.strength()
	}
	return
}

func allExpansionsOf(b bridge, cs []component) (ret []bridge) {
	port := "0"
	for _, c := range b.cs {
		port = c.otherPort(port)
	}
	for _, c := range cs {
		if c.hasPort(port) && !c.in(b) {
			nb := bridge{}
			for _, oc := range b.cs {
				nb.cs = append(nb.cs, oc)
			}
			nb.cs = append(nb.cs, c)
			ret = append(ret, nb)
		}
	}
	return
}

func allBridges(cs []component) (ret []bridge) {
	for _, c := range cs {
		if c.hasPort("0") {
			ret = append(ret, bridge{cs: []component{c}})
		}
	}

	for i := 0; i < len(ret); i++ {
		ae := allExpansionsOf(ret[i], cs)
		ret = append(ret, ae...)
	}
	return
}

func part1() {
	bs := allBridges(allComponents())
	var maxStrength int
	for _, b := range bs {
		if b.strength() > maxStrength {
			maxStrength = b.strength()
		}
	}
	fmt.Println("The max strength of any bridge is:", maxStrength)
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
