package main

import (
	"fmt"
	"time"

	"github.com/gammazero/deque"
)

func playRound(d *deque.Deque) {
	dest := d.Front().(int) - 1
	if dest == 0 {
		dest = 9
	}
	d.Rotate(1)
	a, b, c := d.PopFront(), d.PopFront(), d.PopFront()
	for dest == a || dest == b || dest == c {
		dest--
		if dest == 0 {
			dest = 9
		}
	}
	rot := 1
	for dest != d.Front() {
		rot++
		d.Rotate(1)
	}
	d.Rotate(1)
	d.PushFront(c)
	d.PushFront(b)
	d.PushFront(a)
	d.Rotate(-1 * rot)
}

func finalState(d *deque.Deque) (ret string) {
	for d.Front() != 1 {
		d.Rotate(1)
	}
	d.PopFront()
	for d.Len() > 0 {
		ret += fmt.Sprint(d.PopFront())
	}
	return
}

func part1(in *deque.Deque) {
	for i := 0; i < 100; i++ {
		playRound(in)
	}
	fmt.Println(finalState(in))
}

func part2(in *deque.Deque) {
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	in = input()
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() *deque.Deque {
	ret := deque.Deque{}
	// Example data:
	// for _, i := range []int{3, 8, 9, 1, 2, 5, 4, 6, 7} {
	for _, i := range []int{3, 1, 5, 6, 7, 9, 8, 2, 4} {
		ret.PushBack(i)
	}
	return &ret
}
