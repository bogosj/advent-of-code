package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println(input())
}

func input() [][]string {
	var ret [][]string
	lines := strings.Split(rawinput(), "\n")
	for _, line := range lines {
		vals := strings.Split(line, ")")
		if len(vals) == 2 {
			ret = append(ret, vals)
		}
	}
	return ret
}

func rawinput() string {
	data, _ := ioutil.ReadFile("input.txt")
	return string(data)
}
