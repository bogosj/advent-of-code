package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/2019/day18/vault"
)

func part1() {
	for _, n := range strings.Split("81,132,136", ",") {
		v := vault.New(n + ".txt")
		fmt.Printf("Should be %v: %v\n", n, v.ShortestPath())
	}

	v := vault.New("input.txt")
	fmt.Println(v.ShortestPath())
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
