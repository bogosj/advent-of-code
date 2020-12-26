package main

import (
	"fmt"
	"time"
)

const (
	key1 = 10441485
	key2 = 1004920
	sn   = 7
)

func part1() {
	v, loopSize := 1, 0
	for v != key1 {
		v *= sn
		v %= 20201227
		loopSize++
	}
	v = 1
	for j := 0; j < loopSize; j++ {
		v *= key2
		v %= 20201227
	}
	fmt.Printf("The encryption key is %v\n", v)
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}
