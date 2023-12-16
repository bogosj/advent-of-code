package main

import (
	"errors"
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

func cmp(s1, s2 string) (string, error) {
	if s1 == s2 {
		return "", errors.New("same string")
	}

	var shared string
	var foundDiff bool
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			shared += string(s1[i])
		} else {
			if foundDiff == true {
				return "", errors.New("too many diffs")
			}
			foundDiff = true
		}
	}
	return shared, nil
}

func boxPairs(ss []string) string {
	for _, s1 := range ss {
		for _, s2 := range ss {
			if v, err := cmp(s1, s2); err == nil {
				return v
			}
		}
	}
	return ""
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
	fmt.Println("The shared characters between box pairs is:", boxPairs(input()))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
