package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

func input() []string {
	lines := fileinput.ReadLines("input.txt")
	return lines
}

type chars struct {
	c []rune
}

func (c chars) String() string {
	return string(c.c)
}

func (c *chars) swapPos(from, to string) {
	f := intmath.Atoi(from)
	t := intmath.Atoi(to)
	c.c[f], c.c[t] = c.c[t], c.c[f]
}

func (c *chars) idxOfLetter(l string) int {
	r := rune(l[0])
	for i, v := range c.c {
		if r == v {
			return i
		}
	}
	return -1
}

func (c *chars) swapLetters(a, b string) {
	ia, ib := c.idxOfLetter(a), c.idxOfLetter(b)
	c.c[ia], c.c[ib] = c.c[ib], c.c[ia]
}

func (c *chars) rotate(dir, steps int) {
	if steps == 0 {
		return
	}

	nc := make([]rune, len(c.c))
	for i, v := range c.c {
		ni := ((i + (dir * steps)) + len(c.c)*100) % len(c.c)
		nc[ni] = v
	}
	c.c = nc
}

func (c *chars) rotateBasedOn(ch string) {
	i := c.idxOfLetter(ch)
	rotations := 1 + i
	if i >= 4 {
		rotations++
	}
	c.rotate(1, rotations)
}

func (c *chars) reverse(from, to string) {
	left, right := intmath.Atoi(from), intmath.Atoi(to)
	for left < right {
		c.c[left], c.c[right] = c.c[right], c.c[left]
		left++
		right--
	}
}

func (c *chars) move(from, to string) {
	f, t := intmath.Atoi(from), intmath.Atoi(to)
	dir := 1
	if f > t {
		dir = -1
	}
	for f != t {
		c.swapPos(fmt.Sprintf("%d", f), fmt.Sprintf("%d", f+dir))
		f += dir
	}
}

func scramble(in string) string {
	c := chars{c: []rune(in)}
	for _, inst := range input() {
		f := strings.Fields(inst)
		switch f[0] {
		case "swap":
			switch f[1] {
			case "position":
				c.swapPos(f[2], f[5])
			case "letter":
				c.swapLetters(f[2], f[5])
			}
		case "rotate":
			switch f[1] {
			case "based":
				c.rotateBasedOn(f[6])
			default:
				dir := 1
				if f[1] == "left" {
					dir = -1
				}
				c.rotate(dir, intmath.Atoi(f[2]))
			}
		case "reverse":
			c.reverse(f[2], f[4])
		case "move":
			c.move(f[2], f[5])
		}
	}
	return c.String()
}

func part1() {
	s := "abcdefgh"
	fmt.Printf("The result of scambling %q is %q\n", s, scramble(s))
}

func part2() {
	target := "fbgdceah"
	for perm := range intmath.Permutations([]int{0, 1, 2, 3, 4, 5, 6, 7}) {
		in := make([]rune, 8)
		for i, r := range perm {
			in[i] = rune(r + 'a')
		}
		out := scramble(string(in))
		if out == target {
			fmt.Println(string(in), "scrambles to become", target)
			break
		}
	}
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
