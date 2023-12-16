package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func part1(in map[int][]string) {
	for _, move := range moves() {
		for i := 0; i < move[0]; i++ {
			crate := in[move[1]][0]
			in[move[2]] = append([]string{crate}, in[move[2]]...)
			in[move[1]] = in[move[1]][1:]
		}
	}
	for i := 1; i <= 9; i++ {
		fmt.Print(in[i][0])
	}
	fmt.Println()
}

func part2(in map[int][]string) {
	for _, move := range moves() {
		crates := append([]string{}, in[move[1]][0:move[0]]...)
		in[move[2]] = append(crates, in[move[2]]...)
		in[move[1]] = append([]string{}, in[move[1]][move[0]:]...)
	}
	for i := 1; i <= 9; i++ {
		fmt.Print(in[i][0])
	}
	fmt.Println()
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

func input() map[int][]string {
	return map[int][]string{
		1: {"R", "C", "H"},
		2: {"F", "S", "L", "H", "J", "B"},
		3: {"Q", "T", "J", "H", "D", "M", "R"},
		4: {"J", "B", "Z", "H", "R", "G", "S"},
		5: {"B", "C", "D", "T", "Z", "F", "P", "R"},
		6: {"G", "C", "H", "T"},
		7: {"L", "W", "P", "B", "Z", "V", "N", "S"},
		8: {"C", "G", "Q", "J", "R"},
		9: {"S", "F", "P", "H", "R", "T", "D", "L"},
	}
}

func moves() [][]int {
	ret := [][]int{}
	for _, line := range fileinput.ReadLines("input.txt") {
		if strings.Contains(line, "move") {
			subs := strings.Split(line, " ")
			ret = append(ret, []int{
				intmath.Atoi(subs[1]),
				intmath.Atoi(subs[3]),
				intmath.Atoi(subs[5]),
			})
		}
	}
	return ret
}
