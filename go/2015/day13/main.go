package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/stringperm"
)

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func keyFn(from, to string) string {
	return fmt.Sprintf("%v=>%v", from, to)
}

func input(p string, includeSelf bool) (map[string]int, []string) {
	lines := fileinput.ReadLines("input.txt")
	ret := map[string]int{}
	names := map[string]bool{}
	for _, line := range lines {
		f := strings.Fields(line)
		from := f[0]
		to := f[10][:len(f[10])-1]
		score := atoi(f[3])
		if f[2] == "lose" {
			score *= -1
		}
		ret[keyFn(from, to)] = score
		names[from] = true
	}
	if includeSelf {
		for name := range names {
			ret[keyFn(name, "ME")] = 0
			ret[keyFn("ME", name)] = 0
		}
		names["ME"] = true
	}
	var n []string
	for k := range names {
		n = append(n, k)
	}
	return ret, n
}

func tableScore(scores map[string]int, table []string) (score int) {
	for i := range table {
		p := table[i]
		n1 := (i + 1) % len(table)
		n2 := (len(table) + i - 1) % len(table)
		score += scores[keyFn(p, table[n1])]
		score += scores[keyFn(p, table[n2])]
	}
	return
}

func part1() {
	m, n := input("input.txt", false)
	score := math.MinInt32
	for t := range stringperm.Permutations(n) {
		s := tableScore(m, t)
		if s > score {
			score = s
		}
	}
	fmt.Println("Best score is:", score)
}

func part2() {
	m, n := input("input.txt", true)
	score := math.MinInt32
	for t := range stringperm.Permutations(n) {
		s := tableScore(m, t)
		if s > score {
			score = s
		}
	}
	fmt.Println("Best score (with ME) is:", score)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
