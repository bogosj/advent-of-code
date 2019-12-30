package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	input = "hepxcrrq"
)

func containsThreeStraight(pw string) bool {
	for i := 0; i < len(pw)-2; i++ {
		t := pw[i : i+3]
		if t[1]-t[0] == 1 && t[2]-t[1] == 1 {
			return true
		}
	}
	return false
}

func containsBadLetters(pw string) bool {
	return strings.ContainsAny(pw, "iol")
}

func containsTwoPairs(pw string) bool {
	foundOne := false
	for i := 0; i < len(pw)-1; i++ {
		t := pw[i : i+2]
		if t[0] == t[1] {
			if foundOne {
				return true
			}
			foundOne = true
			i++
		}
	}
	return false
}

func passIsValid(pw string) bool {
	return containsThreeStraight(pw) && containsTwoPairs(pw) && !containsBadLetters(pw)
}

func increment(pw string) string {
	b := []byte(pw)
	idx := len(b) - 1
	for {
		if b[idx] != 'z' {
			b[idx]++
			return string(b)
		}
		b[idx] = 'a'
		idx--
	}
}

func part1() {
	pw := increment(input)
	for !passIsValid(pw) {
		pw = increment(pw)
	}
	fmt.Println("First valid password:", pw)
}

func part2() {
	pw := increment("hepxxyzz")
	for !passIsValid(pw) {
		pw = increment(pw)
	}
	fmt.Println("Second valid password:", pw)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
