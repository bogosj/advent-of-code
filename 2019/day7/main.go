package main

import (
	"fmt"
	"io/ioutil"
	"jamesbogosian.com/advent-of-code/2019/computer"
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
			c := computer.New(input("input.txt"))
			out, err := c.Compute(i)
			if err != nil {
				fmt.Println("err:", err)
			}
			out, err = c.Compute(ampIn)
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
		var computers []*computer.Computer
		ampIn := 0
		for i := 0; i < 5; i++ {
			c := computer.New(input("input.txt"))
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

func input(n string) map[int]int {
	ret := map[int]int{}
	lines := strings.Split(rawinput(n), "\n")
	for i, v := range strings.Split(lines[0], ",") {
		iv, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		ret[i] = iv
	}
	return ret
}

func rawinput(n string) string {
	data, _ := ioutil.ReadFile(n)
	return string(data)
}
