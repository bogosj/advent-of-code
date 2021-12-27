package main

import (
	"fmt"
	"time"
)

var (
	input = `#############
#...........#
###C#D#A#B###
	#B#A#D#C#
	#########`
	sample = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
	#########`
	goal = `#############
#...........#
###A#B#C#D###
	#A#B#C#D#
	#########`
)

func part1() {
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))

	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}
