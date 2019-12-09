package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func part1() {
	img := image{}
	img.parse(input())
	sort.Slice(img.layers, func(i, j int) bool { return img.layers[i].numZeros() < img.layers[j].numZeros() })
	fmt.Println(img.layers[0].digits()[1] * img.layers[0].digits()[2])
}

func part2() {
	img := image{}
	img.parse(input())
	fmt.Println(img.flatten().String())
}

func main() {
	part1()
	part2()
}

func input() []int {
	var ret []int
	line := strings.Split(rawinput(), "\n")
	for _, c := range line[0] {
		ret = append(ret, int(c-48))
	}
	return ret
}

func rawinput() string {
	data, _ := ioutil.ReadFile("input.txt")
	return string(data)
}
