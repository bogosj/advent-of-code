package main

import (
	"fmt"
	"io/ioutil"
	"jamesbogosian.com/advent-of-code/2019/computer"
	"strconv"
	"strings"
	"time"
)

func test() {
	for _, i := range []int{1, 2, 3} {
		c := computer.New(input(fmt.Sprintf("test%v.txt", i)))
		out, err := c.Compute(0)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%v ", out)
		fmt.Println()
	}
}

func part1() {
	c := computer.New(input("input.txt"))
	out, err := c.Compute(1)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("TESTS:", out)
}

func part2() {
	c := computer.New(input("input.txt"))
	out, err := c.Compute(2)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Result:", out)
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
