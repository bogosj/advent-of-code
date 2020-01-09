package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"github.com/bogosj/advent-of-code/intmath"
)

const (
	input = "edjrjqaa"
)

func getMd5(passcode, s string) string {
	h := md5.New()
	io.WriteString(h, passcode+s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

type state struct {
	loc  intmath.Point
	path string
}

func (s state) isValid() bool {
	return s.loc.X >= 0 && s.loc.Y >= 0 && s.loc.X <= 3 && s.loc.Y <= 3
}

func (s state) nextStates(passcode string) (ret []state) {
	hash := getMd5(passcode, s.path)
	var possible []state
	if hash[0] > 'a' {
		ns := s
		ns.loc.Y--
		ns.path += "U"
		possible = append(possible, ns)
	}
	if hash[1] > 'a' {
		ns := s
		ns.loc.Y++
		ns.path += "D"
		possible = append(possible, ns)
	}
	if hash[2] > 'a' {
		ns := s
		ns.loc.X--
		ns.path += "L"
		possible = append(possible, ns)
	}
	if hash[3] > 'a' {
		ns := s
		ns.loc.X++
		ns.path += "R"
		possible = append(possible, ns)
	}
	for _, pos := range possible {
		if pos.isValid() {
			ret = append(ret, pos)
		}
	}
	return
}

func shortestPath(passcode string) string {
	states := []state{{}}
	for len(states) > 0 {
		s := states[0]
		states = states[1:]
		if s.loc.X == 3 && s.loc.Y == 3 {
			return s.path
		}
		states = append(states, s.nextStates(passcode)...)
	}
	return "FAIL"
}

func longestPath(passcode string) (ret int) {
	states := []state{{}}
	for len(states) > 0 {
		s := states[0]
		states = states[1:]
		if s.loc.X == 3 && s.loc.Y == 3 {
			if ret < len(s.path) {
				ret = len(s.path)
			}
		} else {
			states = append(states, s.nextStates(passcode)...)
		}
	}
	return
}

func part1() {
	fmt.Println("The shortest path to the vault is:", shortestPath(input))
}

func part2() {
	fmt.Println("The longest path to the vault is:", longestPath(input))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
