package main

import (
	"fmt"
	"time"
)

const (
	input = 580741
)

type board struct {
	b      []int
	e1, e2 int
}

func newBoard() *board {
	return &board{
		b:  []int{3, 7},
		e2: 1,
	}
}

func (b *board) extend() {
	n := b.b[b.e1] + b.b[b.e2]
	if n > 9 {
		b.b = append(b.b, 1, n%10)
	} else {
		b.b = append(b.b, n)
	}
	b.e1 = (b.e1 + b.b[b.e1] + 1) % len(b.b)
	b.e2 = (b.e2 + b.b[b.e2] + 1) % len(b.b)
}

func scoreAfter(c int) (ret string) {
	b := newBoard()
	for len(b.b) < c+11 {
		b.extend()
	}
	for i := c; i < c+10; i++ {
		ret += fmt.Sprintf("%d", b.b[i])
	}
	return
}

func part1() {
	fmt.Printf("The ten scores after %d recipes are %s\n", input, scoreAfter(input))
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
