package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

const (
	immune = iota
	infection
)

type group struct {
	damage, hp        int
	units, initiative int
	team              int
	attack            string
	immune, weak      []string
}

func input(s string) (ret []*group) {
	atoi := intmath.Atoi
	team := immune
	lines := fileinput.ReadLines(s)
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			i++
			team = infection
			continue
		}
		ng := group{}
		f := strings.Fields(line)
		ng.units = atoi(f[0])
		f = f[4:]
		ng.hp = atoi(f[0])
		f = f[3:]
		for f[0] != "with" {
			if f[0][0] == 'w' || f[0][1] == 'w' {
				f = f[2:]
				for {
					s := f[0]
					s = strings.ReplaceAll(s, ",", "")
					s = strings.ReplaceAll(s, ";", "")
					s = strings.ReplaceAll(s, ")", "")
					ng.weak = append(ng.weak, s)
					f = f[1:]
					if f[0] == "immune" || f[0] == "with" {
						break
					}
				}
			} else {
				f = f[2:]
				for {
					s := f[0]
					s = strings.ReplaceAll(s, ",", "")
					s = strings.ReplaceAll(s, ";", "")
					s = strings.ReplaceAll(s, ")", "")
					ng.immune = append(ng.immune, s)
					f = f[1:]
					if f[0] == "weak" || f[0] == "with" {
						break
					}
				}
			}
		}
		ng.damage = atoi(f[5])
		ng.attack = f[6]
		ng.initiative = atoi(f[10])
		ng.team = team
		ret = append(ret, &ng)
	}
	return
}

func part1() {
	groups := input("input.txt")
	for _, group := range groups {
		fmt.Println(group)
	}
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
