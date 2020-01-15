package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bogosj/advent-of-code/2017/knothash"
	"github.com/bogosj/advent-of-code/intmath"
)

const (
	input = "jzgqcdpd"
)

func allHashes() (ret []string) {
	for i := 0; i < 128; i++ {
		s := fmt.Sprintf("%s-%d", input, i)
		ret = append(ret, knothash.Hash(s))
	}
	return
}

func makeMap(hashes []string) map[intmath.Point]int {
	ret := map[intmath.Point]int{}
	for y, hash := range hashes {
		var x int
		for _, c := range hash {
			if i, err := strconv.ParseInt(string(c), 16, 64); err == nil {
				bin := strconv.FormatInt(i, 2)
				for _, c := range bin {
					ret[intmath.Point{X: x, Y: y}] = int(c - '0')
					x++
				}
			} else {
				panic(err)
			}
		}
	}
	return ret
}

func part1() {
	var sum int
	for _, v := range makeMap(allHashes()) {
		sum += v
	}
	fmt.Println("The number of squares used is:", sum)
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
