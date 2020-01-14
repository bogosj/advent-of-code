package main

import (
	"fmt"
	"time"
)

const (
	ascii = "183,0,31,146,254,240,223,150,2,206,161,1,255,232,199,88"
)

var (
	input = []int{183, 0, 31, 146, 254, 240, 223, 150, 2, 206, 161, 1, 255, 232, 199, 88}
)

func asciiInput() (ret []int) {
	for _, r := range ascii {
		ret = append(ret, int(r))
	}
	ret = append(ret, 17, 31, 73, 47, 23)
	return
}

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

func rotateList2(in []int) []int {
	var idx, skipSize int
	for i := 0; i < 64; i++ {
		for _, length := range asciiInput() {
			reverse(in, idx, length)
			idx = (idx + skipSize + length) % len(in)
			skipSize++
		}
	}
	return in
}

func denseHash(in []int) (ret string) {
	for len(in) > 0 {
		part := in[0:16]
		in = in[16:]
		v := part[0]
		for i := 1; i < len(part); i++ {
			v ^= part[i]
		}
		n := fmt.Sprintf("%x", v)
		if len(n) == 1 {
			n = "0" + n
		}
		ret += n
	}
	return
}

func part1() {
	l := makeList()
	l = rotateList(l)
	fmt.Println("The product of the first two numbers is:", l[0]*l[1])
}

func part2() {
	l := makeList()
	l = rotateList2(l)
	fmt.Println("The Knot Hash is:", denseHash(l))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
