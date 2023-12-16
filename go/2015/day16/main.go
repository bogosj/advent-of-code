package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func rFilter(candidates []string, filter string) (ret []string) {
	f := strings.Fields(filter)
	for _, c := range candidates {
		if !strings.Contains(c, f[0]) || strings.Contains(c, filter) {
			ret = append(ret, c)
		}
	}
	return
}

func rFilter2(candidates []string, filter string) (ret []string) {
	f := strings.Fields(filter)
	for _, c := range candidates {
		if !strings.Contains(c, f[0]) {
			ret = append(ret, c)
			continue
		}
		re := regexp.MustCompile(fmt.Sprintf(`%v (\d+)`, f[0]))
		m := re.FindStringSubmatch(c)
		v := intmath.Atoi(m[1])
		fv := intmath.Atoi(f[1])

		switch f[0] {
		case "cats:", "trees:":
			if v > fv {
				ret = append(ret, c)
			}
		case "pomeranians:", "goldfish:":
			if v < fv {
				ret = append(ret, c)
			}
		default:
			if strings.Contains(c, filter) {
				ret = append(ret, c)
			}
		}
	}
	return
}

func part1() {
	readings := fileinput.ReadLines("reading.txt")
	sues := fileinput.ReadLines("sues.txt")
	for _, reading := range readings {
		sues = rFilter(sues, reading)
	}
	fmt.Println(sues)
}

func part2() {
	readings := fileinput.ReadLines("reading.txt")
	sues := fileinput.ReadLines("sues.txt")
	for _, reading := range readings {
		sues = rFilter2(sues, reading)
	}
	fmt.Println(sues)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
