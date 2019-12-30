package main

import (
	"fmt"
	"math"
	"time"

	"github.com/bogosj/advent-of-code/2019/computer"
	"github.com/bogosj/advent-of-code/intmath"
)

func permutations(in []int) (p [][]int) {
	return intmath.Permutations(in)
}

func part1() {
	maxOutput := math.MinInt64
	allPhases := permutations([]int{0, 1, 2, 3, 4})
	for _, phase := range allPhases {
		ampIn := 0
		for _, i := range phase {
			c := computer.New("input.txt")
			in := make(chan int, 2)
			in <- i
			in <- ampIn
			ampIn = <-c.Compute(in)
		}
		if ampIn > maxOutput {
			maxOutput = ampIn
		}
	}
	fmt.Println(maxOutput)
}

func part2() {
	maxOutput := math.MinInt64
	allPhases := permutations([]int{5, 6, 7, 8, 9})
	for _, phase := range allPhases {
		var computers []*computer.Computer
		ampIn := 0
		for i := 0; i < 5; i++ {
			c := computer.New("input.txt")
			computers = append(computers, c)
		}
		i := 0
		maxEOutput := 0
		for {
			in := make(chan int, 2)
			if len(phase) > 0 {
				in <- phase[0]
				phase = phase[1:]
			}
			in <- ampIn
			c := computers[i%len(computers)]

			out, ok := <-c.Compute(in)
			if !ok {
				break
			}
			if i%len(computers) == 4 {
				if out > maxEOutput {
					maxEOutput = out
				}
			}
			ampIn = out
			i++
		}
		if maxEOutput > maxOutput {
			maxOutput = maxEOutput
		}
	}
	fmt.Println(maxOutput)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
