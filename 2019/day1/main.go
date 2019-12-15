package main

import (
	"jamesbogosian.com/advent-of-code/2019/fileinput"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

func main() {
	totalFuel := 0
	for _, m := range input() {
		totalFuel += fuelRequired(m)
	}
	fmt.Printf("fuel required: %v\n", totalFuel)

	totalFuel = 0
	for _, m := range input() {
		totalFuel += fuelRequiredIfFuelHasMass(m)
	}
	fmt.Printf("fuel required if fuel has mass: %v\n", totalFuel)
}

func input() []int {
	var ret []int
	for _, v := fileinput.ReadLines("input.txt") {
		iv, _ := strconv.Atoi(v)
		ret = append(ret, iv)
	}
	return ret
}
