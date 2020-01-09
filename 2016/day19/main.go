package main

import (
	"fmt"
	"time"
)

const (
	input = 3018458
)

type elf struct {
	next *elf
	id   int
}

func (e *elf) String() string {
	return fmt.Sprintf("Elf ID: %d", e.id)
}

func part1() {
	elves := map[int]*elf{}
	for i := 1; i <= input; i++ {
		e := elf{id: i}
		elves[i] = &e
	}
	for i := 1; i < input; i++ {
		curr := elves[i]
		next := elves[i+1]
		curr.next = next
	}
	first := elves[1]
	last := elves[input]
	last.next = first

	curr := elves[1]
	for curr != curr.next {
		curr.next = curr.next.next
		curr = curr.next
	}
	fmt.Println("Last elf standing is", curr)
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
