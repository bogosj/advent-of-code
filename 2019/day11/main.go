package main

import (
	"fmt"
	"io/ioutil"
	"jamesbogosian.com/advent-of-code/2019/computer"
	"jamesbogosian.com/advent-of-code/2019/day11/robot"
	"strconv"
	"strings"
	"time"
)

func part1() {
	c := computer.New(input("input.txt"))
	r := robot.New(c)
	r.Paint(0)
	panels := r.PrintHull()
	fmt.Println("Number of panels:", panels)
}

func part2() {
	c := computer.New(input("input.txt"))
	r := robot.New(c)
	r.Paint(1)
	r.PrintHull()
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
