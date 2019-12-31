package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func processMap(m map[string]interface{}, f *float64) {
	for _, v := range m {
		processUnknown(v, f)
	}
}

func processSlice(d []interface{}, f *float64) {
	for _, dd := range d {
		processUnknown(dd, f)
	}
}

func processUnknown(d interface{}, f *float64) {
	if s, ok := d.([]interface{}); ok {
		processSlice(s, f)
	}
	if m, ok := d.(map[string]interface{}); ok {
		processMap(m, f)
	}
	if fv, ok := d.(float64); ok {
		*f += fv
	}
}

func part1() {
	lines := fileinput.ReadLines("input.txt")
	var d interface{}
	var f float64
	err := json.Unmarshal([]byte(lines[0]), &d)
	if err != nil {
		panic(err)
	}
	processSlice(d.([]interface{}), &f)
	fmt.Println("Sum of all numbers:", f)
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
