package main

import (
	"fmt"
	"time"
)

var (
	input = []int{1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1}
)

func dragonCurve(a []int, l int) (ret []int) {
	ret = append(ret, a...)
	for len(ret) < l {
		var rev []int
		for i := len(ret) - 1; i >= 0; i-- {
			if ret[i] == 0 {
				rev = append(rev, 1)
			} else {
				rev = append(rev, 0)
			}
		}
		ret = append(ret, 0)
		ret = append(ret, rev...)
	}
	ret = ret[:l]
	return
}

func checksum(c []int) (ret []int) {
	for i := 0; i < len(c)-1; i += 2 {
		if c[i] == c[i+1] {
			ret = append(ret, 1)
		} else {
			ret = append(ret, 0)
		}
	}
	if len(ret)%2 == 1 {
		return ret
	}
	return checksum(ret)
}

func part1() {
	c := dragonCurve(input, 272)
	fmt.Printf("The checksum is: ")
	for _, i := range checksum(c) {
		fmt.Printf("%d", i)
	}
	fmt.Println()
}

func part2() {
	c := dragonCurve(input, 35651584)
	fmt.Printf("The checksum is: ")
	for _, i := range checksum(c) {
		fmt.Printf("%d", i)
	}
	fmt.Println()
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
