package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func part1(in []string) {
	visited := map[intmath.Point]bool{}
	head := intmath.Point{}
	tail := intmath.Point{}
	visited[tail] = true
	for _, line := range in {
		dir := strings.Split(line, " ")[0]
		count := intmath.Atoi(strings.Split(line, " ")[1])
		for step := 0; step < count; step++ {
			prevHead := head
			head = head.Neighbor(rune(dir[0]))
			tailMove := true
			if tail == head {
				tailMove = false
			}
			for _, neighbor := range head.AllNeighbors() {
				if tail == neighbor {
					tailMove = false
					break
				}
			}
			if tailMove {
				tail = prevHead
			}
			visited[tail] = true
		}
	}
	fmt.Printf("The tail visited %d positions\n", len(visited))
}

func part2(in []string) {
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
