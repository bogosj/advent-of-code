package main

import (
	"fmt"
	"time"

	"github.com/gammazero/deque"
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
	var left, right deque.Deque

	for i := 1; i < input/2; i++ {
		left.PushFront(elf{id: i})
	}
	for i := input / 2; i <= input; i++ {
		right.PushFront(elf{id: i})
	}

	for left.Len() > 1 && right.Len() > 1 {
		if left.Len() > right.Len() {
			left.PopFront()
		} else {
			right.PopBack()
		}
		left.PushFront(right.PopBack())
		right.PushFront((left.PopBack()))
	}
	if left.Len() == 1 {
		fmt.Println(left.Front())
	} else {
		fmt.Println(right.Front())
	}
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
