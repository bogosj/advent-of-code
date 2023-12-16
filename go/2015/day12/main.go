package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func processMap(m map[string]interface{}, f *float64, ignoreRed bool) {
	if ignoreRed {
		for _, v := range m {
			if s, ok := v.(string); ok {
				if s == "red" {
					return
				}
			}
		}
	}
	for _, v := range m {
		processUnknown(v, f, ignoreRed)
	}
}

func processSlice(d []interface{}, f *float64, ignoreRed bool) {
	for _, dd := range d {
		processUnknown(dd, f, ignoreRed)
	}
}

func processUnknown(d interface{}, f *float64, ignoreRed bool) {
	if s, ok := d.([]interface{}); ok {
		processSlice(s, f, ignoreRed)
	}
	if m, ok := d.(map[string]interface{}); ok {
		processMap(m, f, ignoreRed)
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
	processSlice(d.([]interface{}), &f, false)
	fmt.Println("Sum of all numbers:", f)
}

func part2() {
	lines := fileinput.ReadLines("input.txt")
	var d interface{}
	var f float64
	err := json.Unmarshal([]byte(lines[0]), &d)
	if err != nil {
		panic(err)
	}
	processSlice(d.([]interface{}), &f, true)
	fmt.Println("Sum of all numbers:", f)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
