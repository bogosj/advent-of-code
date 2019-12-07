package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	maxOutput := math.MinInt64
	allPhases := permutations([]int{0, 1, 2, 3, 4})
	for _, phase := range allPhases {
		ampIn := 0
		for _, i := range phase {
			ampIn = compute(input(), []int{i, ampIn})
		}
		if ampIn > maxOutput {
			maxOutput = ampIn
		}
	}
	fmt.Println(maxOutput)
}

func input() []int {
	var ret []int
	lines := strings.Split(rawinput(), "\n")
	for _, v := range strings.Split(lines[0], ",") {
		iv, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		ret = append(ret, iv)
	}
	return ret
}

func rawinput() string {
	data, _ := ioutil.ReadFile("input.txt")
	return string(data)
}
