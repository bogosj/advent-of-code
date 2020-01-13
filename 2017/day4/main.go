package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func input() (ret [][]string) {
	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, strings.Fields(line))
	}
	return
}

func passIsValid(in []string) bool {
	m := map[string]bool{}
	for _, word := range in {
		if m[word] {
			return false
		}
		m[word] = true
	}
	return true
}

func part1() {
	var valid int
	for _, pass := range input() {
		if passIsValid(pass) {
			valid++
		}
	}
	fmt.Printf("There are %d valid pass phrases\n", valid)
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
