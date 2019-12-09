package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(input())
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
