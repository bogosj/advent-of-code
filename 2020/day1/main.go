package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func part1() {
	in := input()
	for i, v1 := range in {
		for j, v2 := range in {
			if i != j {
				if v1+v2 == 2020 {
					fmt.Printf("Part 1 answer: %v\n", v1*v2)
					return
				}
			}
		}
	}
}

func part2() {
	in := input()
	for i, v1 := range in {
		for j, v2 := range in {
			for k, v3 := range in {
				if i != j {
					if j != k {
						if v1+v2+v3 == 2020 {
							fmt.Printf("Part 2 answer: %v\n", v1*v2*v3)
							return
						}
					}
				}
			}
		}
	}
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []int {
	var ret []int
	lines := fileinput.ReadLines("input.txt")
	for _, v := range lines {
		iv, _ := strconv.Atoi(v)
		ret = append(ret, iv)
	}
	return ret
}
