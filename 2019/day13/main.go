package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"jamesbogosian.com/advent-of-code/2019/computer"
	"jamesbogosian.com/advent-of-code/2019/day13/game"
)

func part1() {
	c := computer.Computer{Prog: input("input.txt")}
	g := game.Game{C: &c}
	bc := g.LoadGrid()
	fmt.Println("Number of blocks:", bc)
}

func part2() {
	c := computer.Computer{Prog: input("input.txt")}
	g := game.Game{C: &c}
	g.Hack()
	g.PlayGame()
	fmt.Println(g.Score)
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
