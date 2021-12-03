package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func part1(in []string) {
	m := map[int]int{}
	for _, line := range in {
		for i := 0; i < 12; i++ {
			m[i] += int(line[i]) - 48
		}
	}
	g := ""
	e := ""
	for i := 0; i < 12; i++ {
		if m[i] > len(in)/2 {
			g += "1"
			e += "0"
		} else {
			g += "0"
			e += "1"
		}
	}
	gi, _ := strconv.ParseInt(g, 2, 64)
	ei, _ := strconv.ParseInt(e, 2, 64)
	fmt.Println("Part 1 answer:", gi*ei)
}

func filterReadings(in []string, most bool) []string {
	for i := 0; i < 12; i++ {
		next := []string{}
		count := 0
		for _, line := range in {
			count += int(line[i]) - 48
		}
		searchBit := '0'
		if count < (len(in)+1)/2 {
			if most {
				searchBit = '0'
			} else {
				searchBit = '1'
			}
		} else {
			if most {
				searchBit = '1'
			} else {
				searchBit = '0'
			}
		}
		for _, line := range in {
			if rune(line[i]) == searchBit {
				next = append(next, line)
			}
		}
		in = next
		if len(in) == 1 {
			return in
		}
	}
	return in
}

func part2(in []string) {
	m := map[int]int{}
	for _, line := range in {
		for i, c := range line {
			m[i] += intmath.Atoi(string(c))
		}
	}

	oxygen := filterReadings(in, true)
	co2 := filterReadings(in, false)

	oReading, _ := strconv.ParseInt(oxygen[0], 2, 64)
	co2Reading, _ := strconv.ParseInt(co2[0], 2, 64)
	fmt.Println("Part 2 answer:", oReading*co2Reading)
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	in = input()
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []string {
	return fileinput.ReadLines("input.txt")
}
