package main

import (
	"fmt"
	"time"
)

const (
	input = "hepxcrrq"
)

func containsThreeStraight(pw string) bool {

	return false
}

func containsBadLetters(pw string) bool {

	return false
}

func containsTwoPairs(pw string) bool {
	return false
}

func passIsValid(pw string) bool {
	return containsThreeStraight(pw) && containsTwoPairs(pw) && !containsBadLetters(pw)
}

func increment(pw string) string {

}

func part1() {
	pw := increment(input)
	for !passIsValid(pw) {
		pw = increment(pw)
	}
	fmt.Println("First valid password:", pw)
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
