package main

import (
	"fmt"
	"sort"
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

func reorder(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func passIsValid(in []string, anagram bool) bool {
	m := map[string]bool{}
	for _, word := range in {
		if anagram {
			word = reorder(word)
		}
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
		if passIsValid(pass, false) {
			valid++
		}
	}
	fmt.Printf("There are %d valid pass phrases\n", valid)
}

func part2() {
	var valid int
	for _, pass := range input() {
		if passIsValid(pass, true) {
			valid++
		}
	}
	fmt.Printf("There are %d valid pass phrases when using anagrams\n", valid)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
