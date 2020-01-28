package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
	"github.com/gammazero/deque"
)

type point = intmath.Point

func input() string {
	s := fileinput.ReadLines("input.txt")[0]
	s = strings.ReplaceAll(s, "^", "")
	return strings.ReplaceAll(s, "$", "")
}

func buildMap(s string) map[point]string {
	ret := map[point]string{}
	stack := deque.Deque{}
	curr := point{}
	for _, c := range s {
		switch c {
		case '^', '$':
		case '(':
			stack.PushBack(curr)
		case '|':
			curr = stack.Back().(point)
		case ')':
			curr = stack.PopBack().(point)
		default:
			path := ret[curr]
			curr = curr.Neighbor(c)
			newPath := path + string(c)
			if currPath, ok := ret[curr]; ok {
				if len(currPath) > len(newPath) {
					ret[curr] = newPath
				}
			} else {
				ret[curr] = newPath
			}
		}
	}
	return ret
}

func maxPath(s string) (max int) {
	m := buildMap(s)
	for _, v := range m {
		if len(v) > max {
			max = len(v)
		}
	}
	return
}

func roomsWith1000Doors(s string) (count int) {
	m := buildMap(s)
	for _, v := range m {
		if len(v) >= 1000 {
			count++
		}
	}
	return
}

func testCases() {
	s := "^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$"
	fmt.Printf("Should be 18: %d\n", maxPath(s))
	s = "^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"
	fmt.Printf("Should be 23: %d\n", maxPath(s))
	s = "^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$"
	fmt.Printf("Should be 31: %d\n", maxPath(s))
}

func part1() {
	fmt.Println("The max path in the facility is:", maxPath(input()))
}

func part2() {
	fmt.Printf("There are %d rooms that have a shortest path going through at least 1000 doors\n", roomsWith1000Doors(input()))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
