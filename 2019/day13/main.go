package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/computer"
	"github.com/bogosj/advent-of-code/2019/day13/game"
)

func part1() {
	c := computer.New("input.txt")
	g := game.New(c)
	bc := g.LoadGrid()
	fmt.Println("Number of blocks:", bc)
}

func part2() {
	c := computer.New("input.txt")
	g := game.New(c)
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
