package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

var (
	basePattern = []int{0, 1, 0, -1}
)

func input() (ret []int) {
	lines := fileinput.ReadLines("input.txt")
	for _, v := range lines[0] {
		ret = append(ret, int(v-48))
	}
	return
}

func fft(in []int) (ret []int) {
	for elt := 0; elt < len(in); elt++ {
		p := pattern(elt, len(in))
		v := 0
		for i, m := range p {
			v += in[i] * m
		}
		ret = append(ret, intmath.Abs(v)%10)
	}
	return ret
}

func firstEight(in []int, phases int) []int {
	for i := 0; i < phases; i++ {
		in = fft(in)
	}
	return in[:8]
}

func messageOffset(in []int) (ret int) {
	for i := 0; i < 7; i++ {
		ret *= 10
		ret += in[i]
	}
	return
}

func pattern(elt, length int) (ret []int) {
end:
	for {
		for _, p := range basePattern {
			for i := 0; i <= elt; i++ {
				ret = append(ret, p)
				if len(ret) > length {
					break end
				}
			}
		}
	}
	ret = ret[1:]
	return
}

func part1() {
	fmt.Println("First eight digits of input:", firstEight(input(), 100))
}

func part2() {
	in := input()
	offset := messageOffset(in)
	fmt.Println(offset)

	var realSignal []int
	for i := 0; i < 10000; i++ {
		realSignal = append(realSignal, in...)
	}
	for i := 0; i < 100; i++ {
		sum := 0
		for i := len(realSignal) - 1; i >= offset-5; i-- {
			sum += realSignal[i]
			sum %= 10
			realSignal[i] = sum
		}
	}
	fmt.Println("Answer:", realSignal[offset:offset+8])
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 complete in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 complete in", time.Since(start))
}
