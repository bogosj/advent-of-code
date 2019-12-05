package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type opCode struct {
	code  int
	modes []int
}

func (o *opCode) Len() int {
	if o.code == 3 || o.code == 4 {
		return 2
	}
	return len(o.modes) + 2
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

func compute(in []int, i int) {
	s := 0
	for {
		op := parseOpCode(in[s])
		vals := []int{}
		if op.code == 1 || op.code == 2 {
			if op.modes[0] == 0 {
				vals = append(vals, in[in[s+1]])
			} else {
				vals = append(vals, in[s+1])
			}
			if op.modes[1] == 0 {
				vals = append(vals, in[in[s+2]])
			} else {
				vals = append(vals, in[s+2])
			}
		}
		if op.code == 1 {
			// Add
			in[in[s+3]] = vals[0] + vals[1]
		} else if op.code == 2 {
			// Multiply
			in[in[s+3]] = vals[0] * vals[1]
		} else if op.code == 3 {
			// Store
			in[in[s+1]] = i
		} else if op.code == 4 {
			// Output
			output := in[s+1]
			if op.modes[0] == 0 {
				output = in[output]
			}
			fmt.Printf("output: %v\n", output)
		} else if op.code == 99 {
			break
		}
		s += op.Len()
	}
}

func main() {
	compute(input(), 1)
}

func input() []int {
	var ret []int
	for _, v := range strings.Split(rawinput(), ",") {
		iv, _ := strconv.Atoi(v)
		ret = append(ret, iv)
	}
	return ret
}

func rawinput() string {
	data, _ := ioutil.ReadFile("input.txt")
	return string(data)
}
