package main

import (
	"fmt"
	"time"
)

type generator struct {
	mul, div, val int
	picky         bool
	pickyFactor   int
}

func (g *generator) next() int {
	g.val = (g.val * g.mul) % g.div
	if g.picky {
		if g.val%g.pickyFactor == 0 {
			return g.val
		}
		return g.next()
	}
	return g.val
}

func gens() (a, b *generator) {
	a = &generator{
		mul: 16807,
		val: 699,
		div: 2147483647,
	}
	b = &generator{
		mul: 48271,
		val: 124,
		div: 2147483647,
	}
	return
}

func part1() {
	var matches int
	genA, genB := gens()
	for i := 0; i < 40000000; i++ {
		if genA.next()&65535 == genB.next()&65535 {
			matches++
		}
	}
	fmt.Printf("There are %d matches\n", matches)
}

func part2() {
	var matches int
	genA, genB := gens()
	genA.picky = true
	genA.pickyFactor = 4
	genB.picky = true
	genB.pickyFactor = 8
	for i := 0; i < 5000000; i++ {
		if genA.next()&65535 == genB.next()&65535 {
			matches++
		}
	}
	fmt.Printf("There are %d picky matches\n", matches)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
