package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func processInput(in []string) ([]bool, map[intmath.Point]bool) {
	algo := []bool{}
	image := map[intmath.Point]bool{}

	for _, c := range in[0] {
		if c == '#' {
			algo = append(algo, true)
		} else {
			algo = append(algo, false)
		}
	}

	for y, line := range in[2:] {
		for x, c := range line {
			if c == '#' {
				image[intmath.Point{X: x + 100, Y: y + 100}] = true
			}
		}
	}

	return algo, image
}

func runAlgo(algo []bool, image map[intmath.Point]bool, p intmath.Point) bool {
	val := 0
	for y := p.Y - 1; y <= p.Y+1; y++ {
		for x := p.X - 1; x <= p.X+1; x++ {
			np := intmath.Point{X: x, Y: y}
			val = val << 1
			if image[np] {
				val = val + 1
			}
		}
	}
	return algo[val]
}

func step(algo []bool, image map[intmath.Point]bool) map[intmath.Point]bool {
	newImage := map[intmath.Point]bool{}

	for y := 0; y <= 300; y++ {
		for x := 0; x <= 300; x++ {
			p := intmath.Point{X: x, Y: y}
			if y == 0 || x == 0 || y == 300 || x == 300 {
				newImage[p] = !image[p]
			} else {
				newImage[p] = runAlgo(algo, image, p)
			}
		}
	}

	return newImage
}

func part1(in []string) {
	algo, image := processInput(in)
	image = step(algo, image)
	image = step(algo, image)
	count := 0
	for _, c := range image {
		if c {
			count++
		}
	}
	fmt.Println("Part 1 answer:", count)
}

func part2(in []string) {
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	in = input()
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []string {
	return fileinput.ReadLines("input.txt")
}
