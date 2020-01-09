package main

import (
	"fmt"
	"time"
)

const (
	input = "^.^^^..^^...^.^..^^^^^.....^...^^^..^^^^.^^.^^^^^^^^.^^.^^^^...^^...^^^^.^.^..^^..^..^.^^.^.^......."
)

func cellsAbove(r []rune, idx int) string {
	var ret []rune
	for i := idx - 1; i <= idx+1; i++ {
		if i == -1 || i == len(r) {
			ret = append(ret, '.')
		} else {
			ret = append(ret, r[i])
		}
	}
	return string(ret)
}

func nextRow(r []rune) (ret []rune) {
	for i := 0; i < len(r); i++ {
		c := cellsAbove(r, i)
		n := '.'
		if c == "^^." || c == ".^^" || c == "^.." || c == "..^" {
			n = '^'
		}
		ret = append(ret, n)
	}
	return
}

func makeMap(s string, size int) (ret [][]rune) {
	ret = append(ret, []rune(s))
	for len(ret) < size {
		ret = append(ret, nextRow(ret[len(ret)-1]))
	}
	return
}

func countSafeTiles(size int) (ret int) {
	m := makeMap(input, size)
	for _, row := range m {
		for _, cell := range row {
			if cell == '.' {
				ret++
			}
		}
	}
	return
}

func part1() {
	size := 40
	fmt.Printf("There are %d safe tiles with size %d.\n", countSafeTiles(size), size)
}

func part2() {
	size := 400000
	fmt.Printf("There are %d safe tiles with size %d.\n", countSafeTiles(size), size)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
