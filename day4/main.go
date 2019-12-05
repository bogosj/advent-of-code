package main

import (
	"fmt"
)

func rightLength(i []int) bool {
	return len(i) == 6
}

func hasDouble(i []int) bool {
	prev := 0
	for _, j := range i {
		if j == prev {
			return true
		}
		prev = j
	}
	return false
}

type Dbl struct {
	v, start, end int
}

func (d *Dbl) justDouble() bool {
	return d.end-d.start == 1
}

func hasSoloDouble(i []int) bool {
	var dbls []Dbl
	for j := 1; j < len(i); j++ {
		if i[j] == i[j-1] {
			d := Dbl{i[j], j - 1, j}
			for j < len(i) && i[j] == i[j-1] {
				d.end = j
				j++
			}
			dbls = append(dbls, d)
		}
	}
	for _, dv := range dbls {
		if dv.justDouble() {
			return true
		}
	}
	return false
}

func onlyIncrease(i []int) bool {
	prev := 0
	for _, j := range i {
		if j < prev {
			return false
		}
		prev = j
	}
	return true
}

func mightBeValid(i int) bool {
	t := explode(i)
	return rightLength(t) && hasDouble(t) && onlyIncrease(t)
}

func mightBeValidPart2(i int) bool {
	t := explode(i)
	return mightBeValid(i) && hasSoloDouble(t)
}

func explode(i int) []int {
	var ret []int
	for {
		d := i % 10
		ret = append([]int{d}, ret...)
		i /= 10
		if i == 0 {
			break
		}
	}
	return ret
}

func main() {
	hasSoloDouble(explode(111322))
	count := 0
	in := input()
	for i := in[0]; i <= in[1]; i++ {
		if mightBeValid(i) {
			count++
		}
	}
	fmt.Printf("valid, part 1: %v\n", count)

	count = 0
	for i := in[0]; i <= in[1]; i++ {
		if mightBeValidPart2(i) {
			count++
		}
	}
	fmt.Printf("valid, part 2: %v\n", count)
}

func input() []int {
	return []int{284639, 748759}
}
