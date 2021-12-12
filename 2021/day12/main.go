package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func makeCaves(in []string) map[string][]string {
	ret := map[string][]string{}
	for _, rooms := range in {
		pair := strings.Split(rooms, "-")
		ret[pair[0]] = append(ret[pair[0]], pair[1])
	}
	return ret
}

func part1(in []string) {
	caves := makeCaves(in)
	fmt.Println(caves)
}

func part2(in []string) {
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	in = input()
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []string {
	return fileinput.ReadLines("input.txt")
}
