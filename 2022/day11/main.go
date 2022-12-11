package main

import (
	"fmt"
	"sort"
	"time"
)

func part1(in []*monkey) {
	for i := 0; i < 20; i++ {
		for _, m := range in {
			for _, item := range m.items {
				m.inspections++
				item = m.op(item)
				item /= 3
				nextMonkey := m.ifFalse
				if item%m.testDivsor == 0 {
					nextMonkey = m.ifTrue
				}
				in[nextMonkey].items = append(in[nextMonkey].items, item)
			}
			m.items = []int{}
		}
	}
	sort.Slice(in, func(i, j int) bool {
		return in[i].inspections > in[j].inspections
	})
	fmt.Printf("The monkey business level is: %d\n", in[0].inspections*in[1].inspections)
}

func part2(in []*monkey) {
	for i := 0; i < 10000; i++ {
		for _, m := range in {
			for _, item := range m.items {
				m.inspections++
				item = m.op(item)
				item %= 9699690
				nextMonkey := m.ifFalse
				if item%m.testDivsor == 0 {
					nextMonkey = m.ifTrue
				}
				in[nextMonkey].items = append(in[nextMonkey].items, item)
			}
			m.items = []int{}
		}
	}
	sort.Slice(in, func(i, j int) bool {
		return in[i].inspections > in[j].inspections
	})
	fmt.Printf("The monkey business level is: %d\n", in[0].inspections*in[1].inspections)
}

func main() {
	in := input()
	start := time.Now()
	part1(in)
	fmt.Println("Part 1 done in", time.Since(start))
	in = input()
	start = time.Now()
	part2(in)
	fmt.Println("Part 2 done in", time.Since(start))
}

type operation func(in int) int

type monkey struct {
	items           []int
	op              operation
	testDivsor      int
	ifTrue, ifFalse int
	inspections     int
}

func input() []*monkey {
	return []*monkey{
		{
			items:      []int{50, 70, 54, 83, 52, 78},
			op:         func(in int) int { return in * 3 },
			testDivsor: 11,
			ifTrue:     2,
			ifFalse:    7,
		},
		{
			items:      []int{71, 52, 58, 60, 71},
			op:         func(in int) int { return in * in },
			testDivsor: 7,
			ifTrue:     0,
			ifFalse:    2,
		},
		{
			items:      []int{66, 56, 56, 94, 60, 86, 73},
			op:         func(in int) int { return in + 1 },
			testDivsor: 3,
			ifTrue:     7,
			ifFalse:    5,
		},
		{
			items:      []int{83, 99},
			op:         func(in int) int { return in + 8 },
			testDivsor: 5,
			ifTrue:     6,
			ifFalse:    4,
		},
		{
			items:      []int{98, 98, 79},
			op:         func(in int) int { return in + 3 },
			testDivsor: 17,
			ifTrue:     1,
			ifFalse:    0,
		},
		{
			items:      []int{76},
			op:         func(in int) int { return in + 4 },
			testDivsor: 13,
			ifTrue:     6,
			ifFalse:    3,
		},
		{
			items:      []int{52, 51, 84, 54},
			op:         func(in int) int { return in * 17 },
			testDivsor: 19,
			ifTrue:     4,
			ifFalse:    1,
		},
		{
			items:      []int{82, 86, 91, 79, 94, 92, 59, 94},
			op:         func(in int) int { return in + 7 },
			testDivsor: 2,
			ifTrue:     5,
			ifFalse:    3,
		},
	}
}
