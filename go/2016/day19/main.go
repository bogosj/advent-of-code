package main

import (
	"fmt"
	"time"

	"github.com/gammazero/deque"
)

const (
	input = 3018458
)

func part1() {
	var elves deque.Deque
	for i := 1; i <= input; i++ {
		elves.PushFront(i)
	}

	for elves.Len() > 1 {
		elves.PushFront(elves.PopBack())
		elves.PopBack()
	}

	fmt.Println("Last elf standing is:", elves.Front())
}

func part2() {
	var left, right deque.Deque

	for i := 1; i < input/2; i++ {
		left.PushFront(i)
	}
	for i := input / 2; i <= input; i++ {
		right.PushFront(i)
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
	fmt.Print("Last elf standing in the new game is: ")
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
