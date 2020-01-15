package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2017/knothash"
)

const (
	ascii = "183,0,31,146,254,240,223,150,2,206,161,1,255,232,199,88"
)

var (
	input = []int{183, 0, 31, 146, 254, 240, 223, 150, 2, 206, 161, 1, 255, 232, 199, 88}
)

func makeList() []int {
	out := make([]int, 256)
	for i := 0; i < 256; i++ {
		out[i] = i
	}
	return out
}

func reverse(in []int, idx, length int) {
	l := len(in)
	for left, right := idx, idx+length-1; left < right; left, right = left+1, right-1 {
		in[left%l], in[right%l] = in[right%l], in[left%l]
	}
}

func rotateList(in []int) []int {
	var idx int
	for skipSize, length := range input {
		reverse(in, idx, length)
		idx = (idx + skipSize + length) % len(in)
	}
	return in
}

func part1() {
	l := makeList()
	l = rotateList(l)
	fmt.Println("The product of the first two numbers is:", l[0]*l[1])
}

func part2() {
	fmt.Println("The Knot Hash is:", knothash.Hash(ascii))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
