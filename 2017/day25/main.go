package main

import (
	"fmt"
	"time"
)

const (
	steps = 12667664
)

const (
	a = iota
	b
	c
	d
	e
	f
)

type machine struct {
	tape       map[int]int
	state, pos int
}

func newMachine() *machine {
	return &machine{
		tape:  map[int]int{},
		state: a,
	}
}

func (m *machine) checksum() (ret int) {
	for _, v := range m.tape {
		ret += v
	}
	return
}

func (m *machine) step() {
	switch m.state {
	case a:
		switch m.tape[m.pos] {
		case 0:
			m.tape[m.pos] = 1
			m.pos++
			m.state = b
		case 1:
			m.tape[m.pos] = 0
			m.pos--
			m.state = c
		}
	case b:
		switch m.tape[m.pos] {
		case 0:
			m.tape[m.pos] = 1
			m.pos--
			m.state = a
		case 1:
			m.tape[m.pos] = 1
			m.pos++
			m.state = d
		}
	case c:
		switch m.tape[m.pos] {
		case 0:
			m.tape[m.pos] = 0
			m.pos--
			m.state = b
		case 1:
			m.tape[m.pos] = 0
			m.pos--
			m.state = e
		}
	case d:
		switch m.tape[m.pos] {
		case 0:
			m.tape[m.pos] = 1
			m.pos++
			m.state = a
		case 1:
			m.tape[m.pos] = 0
			m.pos++
			m.state = b
		}
	case e:
		switch m.tape[m.pos] {
		case 0:
			m.tape[m.pos] = 1
			m.pos--
			m.state = f
		case 1:
			m.tape[m.pos] = 1
			m.pos--
			m.state = c
		}
	case f:
		switch m.tape[m.pos] {
		case 0:
			m.tape[m.pos] = 1
			m.pos++
			m.state = d
		case 1:
			m.tape[m.pos] = 1
			m.pos++
			m.state = a
		}
	}
}

func part1() {
	m := newMachine()
	for i := 0; i < steps; i++ {
		m.step()
	}
	fmt.Println("After running the machine's checksum is:", m.checksum())
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
