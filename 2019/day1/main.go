package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bogosj/advent-of-code/2019/fileinput"
)

func fuelRequired(mass int) int {
	f := mass/3 - 2
	if f < 0 {
		return 0
	}
	return f
}

func fuelRequiredIfFuelHasMass(mass int) int {
	r := fuelRequired(mass)
	if r > 0 {
		return fuelRequiredIfFuelHasMass(r) + r
	}
	return r
}

func part1() {
	totalFuel := 0
	for _, m := range input() {
		totalFuel += fuelRequired(m)
	}
	fmt.Printf("fuel required: %v\n", totalFuel)
}

func part2() {
	totalFuel := 0
	for _, m := range input() {
		totalFuel += fuelRequiredIfFuelHasMass(m)
	}
	fmt.Printf("fuel required if fuel has mass: %v\n", totalFuel)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []int {
	var ret []int
	lines := fileinput.ReadLines("input.txt")
	for _, v := range lines {
		iv, _ := strconv.Atoi(v)
		ret = append(ret, iv)
	}
	return ret
}
