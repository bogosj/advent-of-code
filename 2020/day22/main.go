package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type deck struct {
	cards []int
}

func (d *deck) score() int {
	score := 0
	for i, mult := 0, len(d.cards); mult > 0; i++ {
		score += d.cards[i] * mult
		mult--
	}
	return score
}

func part1(in []*deck) {
	for len(in[0].cards) > 0 && len(in[1].cards) > 0 {
		if in[0].cards[0] > in[1].cards[0] {
			in[0].cards = append(in[0].cards, in[0].cards[0], in[1].cards[0])
		} else {
			in[1].cards = append(in[1].cards, in[1].cards[0], in[0].cards[0])
		}
		in[0].cards = in[0].cards[1:]
		in[1].cards = in[1].cards[1:]
	}

	winner := in[0]
	if len(in[0].cards) == 0 {
		winner = in[1]
	}
	fmt.Printf("The winner's score is %v\n", winner.score())
}

func part2(in []*deck) {
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

func input() []*deck {
	ret := []*deck{}
	ret = append(ret, &deck{})
	for _, line := range fileinput.ReadLines("input.txt") {
		if line == "" {
			ret = append(ret, &deck{})
			continue
		}
		if strings.HasPrefix(line, "Player") {
			continue
		}
		d := ret[len(ret)-1]
		d.cards = append(d.cards, intmath.Atoi(line))
	}

	return ret
}
