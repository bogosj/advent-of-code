package main

import (
	"fmt"
	"sort"
	"time"

	"jamesbogosian.com/advent-of-code/2019/fileinput"
)

func part1() {
	img := image{}
	img.parse(input())
	sort.Slice(img.layers, func(i, j int) bool { return img.layers[i].numZeros() < img.layers[j].numZeros() })
	fmt.Println(img.layers[0].digits()[1] * img.layers[0].digits()[2])
}

func part2() {
	img := image{}
	img.parse(input())
	fmt.Println(img.flatten().String())
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}

func input() []int {
	var ret []int
	lines := fileinput.ReadLines("input.txt")
	for _, c := range lines[0] {
		ret = append(ret, int(c-48))
	}
	return ret
}
