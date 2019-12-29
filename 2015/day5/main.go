package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

/*
It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
*/
var (
	vowels = map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
)

func hasThreeVowels(s string) bool {
	i := 0
	for _, r := range s {
		if vowels[r] == true {
			i++
		}
		if i == 3 {
			return true
		}
	}
	return false
}

func hasDupe(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
	}
	return false
}

func hasBadString(s string) bool {
	for _, ss := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(s, ss) {
			return true
		}
	}
	return false
}

func nice(s string) bool {
	return hasThreeVowels(s) && hasDupe(s) && !hasBadString(s)
}

func pairRepeatsTwice(s string) bool {
	for len(s) > 2 {
		pair := s[0:2]
		if strings.Contains(s[2:], pair) {
			return true
		}
		s = s[1:]
	}
	return false
}

func dupeWithSplit(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i] == s[i-2] {
			return true
		}
	}
	return false
}

func nice2(s string) bool {
	return pairRepeatsTwice(s) && dupeWithSplit(s)
}

func part1() {
	lines := fileinput.ReadLines("input.txt")
	i := 0
	for _, line := range lines {
		if nice(line) {
			i++
		}
	}
	fmt.Println("Nice strings:", i)
}

func part2() {
	lines := fileinput.ReadLines("input.txt")
	i := 0
	for _, line := range lines {
		if nice2(line) {
			i++
		}
	}
	fmt.Println("Nice strings:", i)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
