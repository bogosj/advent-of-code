package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func part1(in []string) {
	count := 0
	for _, line := range in {
		f := strings.Split(line, " | ")
		for _, digit := range strings.Fields(f[1]) {
			if len(digit) != 5 && len(digit) != 6 {
				count++
			}
		}
	}
	fmt.Println("Part 1 answer:", count)
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func selector(wires []string, length int) (ret []string) {
	for _, w := range wires {
		if len(w) == length {
			ret = append(ret, w)
		}
	}
	return
}

func finder(s string, possibles []string, inCommon int) string {
	for _, os := range possibles {
		common := map[rune]int{}
		for _, c := range s {
			common[c]++
		}
		for _, c := range os {
			common[c]++
		}
		count := 0
		for _, v := range common {
			if v > 1 {
				count++
			}
		}
		if count == inCommon {
			return os
		}
	}
	return ""
}

func findThird(values map[string]int, wires []string, v int) {
	for _, wire := range wires {
		_, ok := values[wire]
		if !ok {
			values[wire] = v
		}
	}
}

func findDigits(wires []string) map[string]int {
	numToString := map[int]string{}
	numToString[1] = selector(wires, 2)[0]
	numToString[7] = selector(wires, 3)[0]
	numToString[4] = selector(wires, 4)[0]
	numToString[8] = selector(wires, 7)[0]
	numToString[2] = finder(numToString[4], selector(wires, 5), 2)
	numToString[3] = finder(numToString[1], selector(wires, 5), 2)
	numToString[6] = finder(numToString[1], selector(wires, 6), 1)
	numToString[9] = finder(numToString[4], selector(wires, 6), 4)

	ret := map[string]int{}
	for k, v := range numToString {
		ret[v] = k
	}
	findThird(ret, selector(wires, 5), 5)
	findThird(ret, selector(wires, 6), 0)
	return ret
}

func decodeLine(line string) int {
	f := strings.Split(line, " | ")
	wires := strings.Fields(f[0])
	display := strings.Fields(f[1])
	for i := 0; i < len(wires); i++ {
		wires[i] = sortString(wires[i])
	}
	for i := 0; i < len(display); i++ {
		display[i] = sortString(display[i])
	}
	digits := findDigits(wires)
	sum := 0
	for _, digit := range display {
		sum *= 10
		sum += digits[digit]
	}
	return sum
}

func part2(in []string) {
	sum := 0
	for _, line := range in {
		sum += decodeLine(line)
	}
	fmt.Println("Part 2 answer:", sum)
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
