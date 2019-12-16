package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/bogosj/advent-of-code/2019/day14/factory"
)

func part1() {
	f := factory.New("input.txt")
	start := time.Now()
	fmt.Println(f.Ore(1))
	fmt.Println("took:", time.Since(start))
}

func part2() {
	f := factory.New("input.txt")
	start := time.Now()
	ore := 1000000000000
	fuel := sort.Search(ore, func(n int) bool {
		return f.Ore(n) > ore
	}) - 1
	fmt.Printf("Can make %d fuel with %d ore\n", fuel, ore)
	fmt.Println("took:", time.Since(start))
}

func main() {
	part1()
	part2()
}
