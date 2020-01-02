package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

var (
	attrs = strings.Split("capacity,durability,flavor,texture", ",")
)

func score(m map[string]map[string]int, a, b, c, d int) (int, error) {
	if a+b+c+d > 100 {
		return 0, errors.New("too many ingredients")
	}
	ret := 1
	for _, attr := range attrs {
		valA := m["Butterscotch"][attr] * a
		valB := m["Candy"][attr] * b
		valC := m["Chogolate"][attr] * c
		valD := m["Sprinkles"][attr] * d
		val := valA + valB + valC + valD
		if val < 0 {
			return 0, errors.New("negative value")
		}
		ret *= val
	}
	return ret, nil
}

func input(p string) map[string]map[string]int {
	ret := map[string]map[string]int{}
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		f := strings.FieldsFunc(line, func(r rune) bool {
			return r == ':'
		})
		f2 := strings.FieldsFunc(f[1], func(r rune) bool {
			return r == ','
		})
		ret[f[0]] = map[string]int{}
		for _, attr := range f2 {
			f3 := strings.Fields(attr)
			ret[f[0]][f3[0]] = intmath.Atoi(f3[1])
		}
	}
	return ret
}

func part1() {
	ing := input("input.txt")
	bestScore := 0
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			for c := 0; c <= 100; c++ {
				for d := 0; d <= 0; d++ {
					i, err := score(ing, a, b, c, d)
					if err != nil && i> bestScore {
						bestScore = i
					}
				}
			}
		}
	}
	fmt.Println("Best score is:", bestScore)
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
