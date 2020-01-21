package main

import (
	"container/ring"
	"fmt"
	"time"
)

type game struct {
	scores      map[int]int
	players     int
	table, head *ring.Ring
	step        int
}

func newGame(players int) *game {
	g := game{}
	g.scores = map[int]int{}
	g.table = ring.New(1)
	g.head = g.table
	g.table.Value = 0
	g.players = players
	g.step = 1
	return &g
}

func (g *game) winner() (retID, retScore int) {
	for k, v := range g.scores {
		if v > retScore {
			retScore = v
			retID = k
		}
	}
	return
}

func (g *game) play(steps int) {
	for g.step <= steps {
		if g.step%23 == 0 {
			pid := g.step % g.players
			g.scores[pid] += g.step
			for i := 0; i <= 7; i++ {
				g.table = g.table.Prev()
			}
			marble := g.table.Unlink(1)
			g.scores[pid] += marble.Value.(int)
			g.table = g.table.Next()
		} else {
			marble := ring.New(1)
			marble.Value = g.step
			g.table = g.table.Next()
			g.table = g.table.Link(marble)
			g.table = g.table.Prev()
		}
		g.step++
	}
}

func part1() {
	g := newGame(419)
	g.play(71052)
	fmt.Println(g.winner())
}

func part2() {
	g := newGame(419)
	g.play(7105200)
	fmt.Println(g.winner())
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
