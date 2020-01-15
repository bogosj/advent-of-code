package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() map[int]int {
	ret := map[int]int{}
	for _, line := range fileinput.ReadLines("input.txt") {
		f := strings.FieldsFunc(line, func(r rune) bool { return r == ',' || r == ' ' })
		f[0] = strings.ReplaceAll(f[0], ":", "")
		ret[intmath.Atoi(f[0])] = intmath.Atoi(f[1])
	}
	return ret
}

func traverse(delay int) (ret int, retDetected bool) {
	for depth, size := range input() {
		period := size*2 - 2
		if (depth+delay)%period == 0 {
			ret += depth * size
			retDetected = true
		}
	}
	return
}

func part1() {
	score, _ := traverse(0)
	fmt.Println("The severity of traversing the firewall is:", score)
}

func part2() {
	for delay := 1; ; delay++ {
		_, detected := traverse(delay)
		if !detected {
			fmt.Printf("A delay of %d picoseconds prevents detection\n", delay)
			return
		}
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
