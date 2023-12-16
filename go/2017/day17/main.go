package main

import (
	"fmt"
	"time"
)

type node struct {
	val  int
	next *node
}

type spinLock struct {
	head, curr *node
	step       int
}

func newSpinLock() *spinLock {
	n := node{val: 0}
	n.next = &n
	return &spinLock{
		head: &n,
		curr: &n,
		step: 301,
	}
}

func (s *spinLock) String() (ret string) {
	var i int
	n := s.head
	for {
		v := fmt.Sprintf("%d", n.val)
		if n == s.curr {
			v = "(" + v + ")"
		}
		ret += v
		n = n.next
		if n == s.head {
			return
		}
		i++
		ret += " "
	}
}

func (s *spinLock) insert(i int) {
	ip := s.curr
	for i := 0; i < s.step; i++ {
		ip = ip.next
	}
	n := &node{
		val:  i,
		next: ip.next,
	}
	s.curr = n
	ip.next = n
}

func (s *spinLock) valueAfter(i int) int {
	n := s.head
	for {
		if n.val == i {
			return n.next.val
		}
		n = n.next
	}
}

func part1() {
	s := newSpinLock()
	for i := 1; i <= 2017; i++ {
		s.insert(i)
	}
	fmt.Printf("The value after %d is %d\n", 2017, s.valueAfter(2017))
}

func part2() {
	s := newSpinLock()
	for i := 1; i <= 50000000; i++ {
		s.insert(i)
	}
	fmt.Printf("The value after %d is %d\n", 0, s.valueAfter(0))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
