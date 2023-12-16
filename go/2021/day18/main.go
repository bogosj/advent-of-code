package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

var (
	pairFinder = regexp.MustCompile(`\[\d+,\d+\]`)
)

func findStartOfPair(in string) int {
	startOfPair := -1
	depth := 0
	for i, c := range in {
		if c == '[' {
			depth++
		}
		if c == ']' {
			depth--
		}
		if depth == 5 {
			startOfPair = i
			break
		}
	}
	return startOfPair
}

func extractPair(in string, startOfPair int) (int, int, string) {
	splitVal := strings.SplitN(in[startOfPair+1:], "]", 2)
	numbers := strings.Split(splitVal[0], ",")
	out := fmt.Sprintf("%sx%s", in[:startOfPair], splitVal[1])
	return intmath.Atoi(numbers[0]), intmath.Atoi(numbers[1]), out
}

func explode(in string) string {
	startOfPair := findStartOfPair(in)
	if startOfPair == -1 {
		return in
	}

	p1, p2, inX := extractPair(in, startOfPair)
	splitVal := strings.SplitN(inX, "x", 2)
	// Add p1 to left

	leftEnd := strings.LastIndexAny(splitVal[0], "0123456789")
	leftStart := leftEnd
	if leftStart >= 0 {
		for splitVal[0][leftStart] >= '0' && splitVal[0][leftStart] <= '9' {
			leftStart--
		}
		leftStart++
		leftEnd++
		val := intmath.Atoi(splitVal[0][leftStart:leftEnd]) + p1
		splitVal[0] = fmt.Sprintf("%s%d%s", splitVal[0][:leftStart], val, splitVal[0][leftEnd:])
	}

	// Add p2 to right
	rightStart := strings.IndexAny(splitVal[1], "0123456789")
	if rightStart >= 0 {
		rightEnd := strings.IndexAny(splitVal[1][rightStart+1:], ",[]") + rightStart + 1
		val := intmath.Atoi(splitVal[1][rightStart:rightEnd]) + p2
		splitVal[1] = fmt.Sprintf("%s%d%s", splitVal[1][:rightStart], val, splitVal[1][rightEnd:])
	}
	// Replace x with 0
	inX = strings.Join(splitVal, "x")
	return strings.Replace(inX, "x", "0", 1)
}

func split(in string) string {
	// Find first two digit number, replace with a pair
	start := -1
	for i, c := range in {
		if c >= '0' && c <= '9' {
			if in[i+1] >= '0' && in[i+1] <= '9' {
				start = i
				break
			}
		}
	}
	if start >= 0 {
		val := intmath.Atoi(in[start : start+2])
		left, right := val/2, val/2
		if val%2 == 1 {
			right++
		}
		return fmt.Sprintf("%s[%d,%d]%s", in[:start], left, right, in[start+2:])
	} else {
		return in
	}
}

func add(in1, in2 string) string {
	val := fmt.Sprintf("[%s,%s]", in1, in2)
	for {
		val2 := explode(val)
		if val != val2 {
			val = val2
			continue
		}
		val2 = split(val)
		if val != val2 {
			val = val2
			continue
		}
		break
	}
	return val
}

func magnitude(in string) int {
	pairs := pairFinder.FindAllString(in, -1)
	for len(pairs) > 0 {
		for _, pair := range pairs {
			vals := pair[1 : len(pair)-1]
			p1 := intmath.Atoi(strings.Split(vals, ",")[0])
			p2 := intmath.Atoi(strings.Split(vals, ",")[1])
			val := p1*3 + p2*2
			in = strings.ReplaceAll(in, pair, fmt.Sprint(val))
		}
		pairs = pairFinder.FindAllString(in, -1)
	}
	return intmath.Atoi(in)
}

func part1(in []string) {
	problem := in[0]
	for i := 1; i < len(in); i++ {
		problem = add(problem, in[i])
	}
	fmt.Println("Part 1 answer:", magnitude(problem))
}

func part2(in []string) {
	mags := []int{}
	for l := 0; l < len(in); l++ {
		for r := 0; r < len(in); r++ {
			if l == r {
				continue
			}
			mags = append(mags, magnitude(add(in[l], in[r])))
		}
	}
	fmt.Println("Part 2 answer:", intmath.Max(mags...))
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

func input() []string {
	return fileinput.ReadLines("input.txt")
}
