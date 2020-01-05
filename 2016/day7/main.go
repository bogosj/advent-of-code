package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

var (
	re = regexp.MustCompile(`\[(.*?)\]`)
)

func input() []string {
	lines := fileinput.ReadLines("input.txt")
	return lines
}

func containsABBA(s string) bool {
	for len(s)>=4 {
		if s[0]==s[3] && s[1]==s[2] && s[0]!=s[1]{
			return true
		}
		s=s[1:]
	}
	return false
}

func supportsTLS(s string) bool {
	if !containsABBA(re.ReplaceAllString(s, "*")) {
		return false
	}
	for _, match:=range re.FindAllStringSubmatch(s, -1) {
		if containsABBA(match[1]) {
			return false
		}
	}
	return true
}

func part1() {
	in := input()
	c:=0
	for _, line:=range in {
		if supportsTLS(line) {
			c++
		}
	}
	fmt.Println(c, "IPs support TLS")
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
