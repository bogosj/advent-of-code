package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func firstWinningBoard(nums []int, boards []board) ([]int, board) {
	n := 5
	for n <= len(nums) {
		for _, b := range boards {
			if b.winning(nums[:n]) {
				return nums[:n], b
			}
		}
		n++
	}
	return nil, board{}
}

func part1(nums []int, boards []board) {
	calls, b := firstWinningBoard(nums, boards)
	fmt.Println("Part 1 answer:", b.score(calls))
}

func part2(nums []int, boards []board) {
}

func main() {
	nums, boards := input()
	start := time.Now()
	part1(nums, boards)
	fmt.Println("Part 1 done in", time.Since(start))
	nums, boards = input()
	start = time.Now()
	part2(nums, boards)
	fmt.Println("Part 2 done in", time.Since(start))
}

type board struct {
	b [][]int
}

func inArray(num int, nums []int) bool {
	for _, n := range nums {
		if num == n {
			return true
		}
	}
	return false
}

func (b board) score(nums []int) int {
	sum := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !inArray(b.b[x][y], nums) {
				sum += b.b[x][y]
			}
		}
	}
	return sum * nums[len(nums)-1]
}

func (b board) winning(nums []int) bool {
	for x := 0; x < 5; x++ {
		if inArray(b.b[x][0], nums) &&
			inArray(b.b[x][1], nums) &&
			inArray(b.b[x][2], nums) &&
			inArray(b.b[x][3], nums) &&
			inArray(b.b[x][4], nums) {
			return true
		}
	}
	for y := 0; y < 5; y++ {
		if inArray(b.b[0][y], nums) &&
			inArray(b.b[1][y], nums) &&
			inArray(b.b[2][y], nums) &&
			inArray(b.b[3][y], nums) &&
			inArray(b.b[4][y], nums) {
			return true
		}
	}
	return false
}

func input() ([]int, []board) {
	f := fileinput.ReadLines("input.txt")

	nums := []int{}
	for _, num := range strings.Split(f[0], ",") {
		nums = append(nums, intmath.Atoi(num))
	}

	boards := []board{}
	var b board
	for _, line := range f[2:] {
		if line == "" {
			boards = append(boards, b)
			b = board{}
			b.b = [][]int{}
		} else {
			row := []int{}
			for _, n := range strings.Fields(line) {
				row = append(row, intmath.Atoi(n))
			}
			b.b = append(b.b, row)
		}
	}

	return nums, boards
}
