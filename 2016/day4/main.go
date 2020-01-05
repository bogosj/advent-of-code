package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type room struct {
	name, checksum string
	sector         int
}

func (r *room) isValid() bool {
	checksum := map[rune]int{}
	for _, r := range r.name {
		checksum[r]++
	}
	cs := ""
	for i := len(r.name); i > 0; i-- {
		for ch := 'a'; ch <= 'z'; ch++ {
			if checksum[ch] == i {
				cs += string(ch)
				if len(cs) == 5 {
					return cs == r.checksum
				}
			}
		}
	}
	return false
}

func (r *room) decrypted() (ret string) {
	cycle := r.sector % 26
	for _, ch := range r.name {
		if ch == '-' {
			ret += " "
			continue
		}
		nc := int(ch) + cycle
		if nc > 'z' {
			nc -= 26
		}
		ret += string(rune(nc))
	}
	ret += fmt.Sprintf(" - %d", r.sector)
	return
}

func input() (ret []room) {
	lines := fileinput.ReadLines("input.txt")
	for _, line := range lines {
		f := strings.FieldsFunc(line, func(r rune) bool {
			return r == '['
		})
		f2 := strings.FieldsFunc(f[0], func(r rune) bool {
			return r == '-'
		})
		name := strings.Join(f2[:len(f2)-1], "-")
		ret = append(ret, room{name: name, sector: intmath.Atoi(f2[len(f2)-1]), checksum: f[1][:5]})
	}
	return
}

func part1() {
	sum := 0
	for _, r := range input() {
		if r.isValid() {
			sum += r.sector
		}
	}
	fmt.Println("Sum of valid sectors:", sum)
}

func part2() {
	for _, r := range input() {
		fmt.Println(r.decrypted())
	}
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
