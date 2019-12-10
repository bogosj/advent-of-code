package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func test() {
	for _, i := range []int{1, 2, 3} {
		c := newComputer()
		c.prog = input(fmt.Sprintf("test%v.txt", i))
		out := c.compute(nil)
		fmt.Printf("%v ", out)
		fmt.Println()
	}
}

func part1() {
	c := newComputer()
	c.prog = input("input.txt")
	output := c.compute([]int{1})
	fmt.Println("TESTS:", output)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
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
