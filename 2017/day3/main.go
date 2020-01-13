package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/intmath"
)

const (
	input = 277678
)

func makeMem() map[intmath.Point]int {
	ret := map[intmath.Point]int{}
	i := 1
	var cur, next intmath.Point
	for {
		//move down until nothight right
		for {
			ret[cur] = i
			i++
			if i > input {
				return ret
			}
			next = cur
			next.X++
			_, ok := ret[next]
			if !ok {
				break
			}
			cur.Y++
		}
		cur = next
		//move right until nothing above
		for {
			ret[cur] = i
			i++
			if i > input {
				return ret
			}
			next = cur
			next.Y--
			_, ok := ret[next]
			if !ok {
				break
			}
			cur.X++
		}
		cur = next
		//move up until nothing left
		for {
			ret[cur] = i
			i++
			if i > input {
				return ret
			}
			next = cur
			next.X--
			_, ok := ret[next]
			if !ok {
				break
			}
			cur.Y--
		}
		cur = next
		//move left until nothing below
		for {
			ret[cur] = i
			i++
			if i > input {
				return ret
			}
			next = cur
			next.Y++
			_, ok := ret[next]
			if !ok {
				break
			}
			cur.X--
		}
		cur = next
	}
}

func sumNeighbors(p intmath.Point, mem map[intmath.Point]int) (ret int) {
	for _, n := range p.AllNeighbors() {
		if v, ok := mem[n]; ok {
			ret += v
		}
	}
	if ret == 0 {
		ret = 1
	}
	return
}

func makeMemSquares() int {
	mem := map[intmath.Point]int{}
	var cur, next intmath.Point
	for {
		//move down until nothight right
		for {
			mem[cur] = sumNeighbors(cur, mem)
			if mem[cur] > input {
				return mem[cur]
			}
			next = cur
			next.X++
			_, ok := mem[next]
			if !ok {
				break
			}
			cur.Y++
		}
		cur = next
		//move right until nothing above
		for {
			mem[cur] = sumNeighbors(cur, mem)
			if mem[cur] > input {
				return mem[cur]
			}
			next = cur
			next.Y--
			_, ok := mem[next]
			if !ok {
				break
			}
			cur.X++
		}
		cur = next
		//move up until nothing left
		for {
			mem[cur] = sumNeighbors(cur, mem)
			if mem[cur] > input {
				return mem[cur]
			}
			next = cur
			next.X--
			_, ok := mem[next]
			if !ok {
				break
			}
			cur.Y--
		}
		cur = next
		//move left until nothing below
		for {
			mem[cur] = sumNeighbors(cur, mem)
			if mem[cur] > input {
				return mem[cur]
			}
			next = cur
			next.Y++
			_, ok := mem[next]
			if !ok {
				break
			}
			cur.X--
		}
		cur = next
	}
}

func part1() {
	for k, v := range makeMem() {
		if v == input {
			fmt.Println(k.ManhattanDistanceTo(intmath.Point{}), "steps are required")
		}
	}
}

func part2() {
	fmt.Println(makeMemSquares(), "is the first value larger than the input")
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
