package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func part1() {
	maxOutput := math.MinInt64
	allPhases := permutations([]int{0, 1, 2, 3, 4})
	for _, phase := range allPhases {
		ampIn := 0
		for _, i := range phase {
			c := computer{prog: input()}
			out, err := c.compute([]int{i, ampIn})
			if err != nil {
				fmt.Println("err:", err)
			}
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
		computers := make([]computer, 5)
		ampIn := 0
		for i := 0; i < len(computers); i++ {
			computers[i].prog = input()
		}
		i := 0
		maxEOutput := 0
		for {
			in := []int{ampIn}
			if len(phase) > 0 {
				in = append([]int{phase[0]}, in...)
				phase = phase[1:]
			}
			out, err := computers[i%len(computers)].compute(in)
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

func input() []int {
	var ret []int
	lines := strings.Split(rawinput(), "\n")
	for _, v := range strings.Split(lines[0], ",") {
		iv, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		ret = append(ret, iv)
	}
	return ret
}

func rawinput() string {
	data, _ := ioutil.ReadFile("input.txt")
	return string(data)
}
