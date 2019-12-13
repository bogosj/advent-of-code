package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func part1() {
	c := computer{prog: input("input.txt")}
	g := game{c: &c}
	g.loadGrid()
	fmt.Println("Number of blocks:", g.blockCount())
}

func part2() {
	c := computer{prog: input("input.txt")}
	g := game{c: &c}
	g.hack()
	g.playGame()
	fmt.Println(g.score)
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
