package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type opCode struct {
	code  int
	modes []int
}

func (o *opCode) Len() int {
	if o.code == 3 || o.code == 4 {
		return 2
	}
	if o.code == 5 || o.code == 6 {
		return 3
	}
	return len(o.modes) + 2
}

func (o *opCode) ReadVals() bool {
	return o.code == 1 || o.code == 2 || (o.code >= 5 && o.code <= 8)
}

func parseOpCode(i int) opCode {
	ret := opCode{code: i % 100}
	i /= 100
	for i != 0 {
		ret.modes = append(ret.modes, i%10)
		i /= 10
	}
	for len(ret.modes) < 2 {
		ret.modes = append(ret.modes, 0)
	}
	return ret
}

func compute(prog []int, i int) {
	pc := 0
Loop:
	for {
		op := parseOpCode(prog[pc])
		vals := []int{}
		if op.ReadVals() {
			if op.modes[0] == 0 {
				vals = append(vals, prog[prog[pc+1]])
			} else {
				vals = append(vals, prog[pc+1])
			}
			if op.modes[1] == 0 {
				vals = append(vals, prog[prog[pc+2]])
			} else {
				vals = append(vals, prog[pc+2])
			}
		}
		switch op.code {
		case 1: // Add
			prog[prog[pc+3]] = vals[0] + vals[1]
		case 2: // Multiply
			prog[prog[pc+3]] = vals[0] * vals[1]
		case 3: // Store
			prog[prog[pc+1]] = i
		case 4: // Output
			output := prog[pc+1]
			if op.modes[0] == 0 {
				output = prog[output]
			}
			fmt.Printf("output: %v\n", output)
		case 5:
			if vals[0] != 0 {
				pc = vals[1]
				continue
			}
		case 6:
			if vals[0] == 0 {
				pc = vals[1]
				continue
			}
		case 7:
			if vals[0] < vals[1] {
				prog[prog[pc+3]] = 1
			} else {
				prog[prog[pc+3]] = 0
			}
		case 8:
			if vals[0] == vals[1] {
				prog[prog[pc+3]] = 1
			} else {
				prog[prog[pc+3]] = 0
			}
		case 99:
			break Loop
		}
		pc += op.Len()
	}
}

func main() {
	start := time.Now()
	compute(input(), 1)
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	compute(input(), 5)
	fmt.Println("Part 2 done in:", time.Since(start))
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