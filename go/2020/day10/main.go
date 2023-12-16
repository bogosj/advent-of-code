package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func part1(in []int) (ones, threes int) {
	for i := 0; i < len(in)-1; i++ {
		switch in[i+1] - in[i] {
		case 1:
			ones++
		case 3:
			threes++
		}
	}
	fmt.Printf("The ones * threes = %v\n", ones*threes)
	return
}

var (
	results map[string]int
)

func keyFromSlice(in []int) string {
	s := []string{}
	for _, i := range in {
		s = append(s, fmt.Sprintf("%v", i))
	}
	return strings.Join(s, ",")
}

func combosFrom(in []int) (ret int) {
	if len(in) == 1 {
		return
	}
	if len(in) == 2 {
		return 1
	}
	for i := 1; i < len(in) && in[i]-in[0] <= 3; i++ {
		v, ok := results[keyFromSlice(in[i:])]
		if !ok {
			v = combosFrom(in[i:])
			results[keyFromSlice(in[i:])] = v
		}
		ret += v
	}
	return
}

func part2(in []int) {
	results = map[string]int{}
	fmt.Printf("There are %v possible combinations.\n", combosFrom(in))
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

func input() []int {
	ret := []int{}

	for _, line := range fileinput.ReadLines("input.txt") {
		ret = append(ret, intmath.Atoi(line))
	}

	ret = append(ret, 0)
	ret = append(ret, intmath.Max(ret...)+3)
	sort.Ints(ret)

	return ret
}
