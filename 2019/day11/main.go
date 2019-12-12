package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func part1() {
	c := computer{}
	c.prog = input("input.txt")

	r := robot{}
	r.c = &c
	r.paint(0)
	panels := r.printHull()
	fmt.Println("Number of panels:", panels)
}

func part2() {
	c := computer{}
	c.prog = input("input.txt")

	r := robot{}
	r.c = &c
	r.paint(1)
	r.printHull()
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
