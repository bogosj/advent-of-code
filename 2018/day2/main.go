package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func input() []string {
	return fileinput.ReadLines("input.txt")
}

func checksum(s string) (two, three int) {
	count := map[rune]int{}
	for _, c := range s {
		count[c]++
	}
	for _, v := range count {
		switch v {
		case 2:
			two = 1
		case 3:
			three = 1
		}
	}
	return
}

func part1() {
	var twos, threes int
	for _, s := range input() {
		two, three := checksum(s)
		twos += two
		threes += three
	}
	fmt.Println("The checksum is:", twos*threes)
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
