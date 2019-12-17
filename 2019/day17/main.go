package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/2019/computer"
	"github.com/bogosj/advent-of-code/2019/day17/camera"
)

func part1() {
	com := computer.New("input.txt")
	c := camera.New(com)
	c.CaptureImage()
	sum := 0
	for _, in := range c.Intersections() {
		sum += in.X * in.Y
	}
	fmt.Println("Alignment:", sum)
}

func part2() {
	com := computer.New("input.txt")
	c := camera.New(com)
	fmt.Println("Dust cleaned:", c.Notify())
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 complete in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 complete in", time.Since(start))
}
