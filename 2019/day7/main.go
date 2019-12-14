package main

import (
	"fmt"
	"math"
	"time"

	"jamesbogosian.com/advent-of-code/2019/computer"
)

func permutations(in []int) (p [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			p = append(p, append([]int{}, a...))
		} else {
			for i := k; i < len(in); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(in, 0)
	return p
}

func part1() {
	maxOutput := math.MinInt64
	allPhases := permutations([]int{0, 1, 2, 3, 4})
	for _, phase := range allPhases {
		ampIn := 0
		for _, i := range phase {
			c := computer.New("input.txt")
			out, _ := c.Compute(i)
			out, _ = c.Compute(ampIn)
			ampIn = out
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
			_, err := c.Compute(phase[i])
			if err != nil {
				fmt.Println("error:", err)
			}
			computers = append(computers, c)
		}
		i := 0
		maxEOutput := 0
		for {
			out, err := computers[i%len(computers)].Compute(ampIn)
			if err != nil {
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
