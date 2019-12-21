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
	fmt.Println("Part 1 answer:", v.ShortestPath())
}

func part2() {
	var ans int
	for _, n := range strings.Split("1,2,3,4", ",") {
		v := vault.New(fmt.Sprintf("part2_v%v.txt", n))
		ans += v.ShortestPath()
	}
	fmt.Println("Part 2 answer:", ans)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}
