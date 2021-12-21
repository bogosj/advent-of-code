package main

import (
	"fmt"
	"time"

	"github.com/bogosj/advent-of-code/intmath"
)

/*
Player 1 starting position: 7
Player 2 starting position: 10
*/

type player struct {
	position, score int
}

func (p player) String() string {
	return fmt.Sprintf("Position: %d, Score: %d", p.position, p.score)
}

// takeTurn returns the state of the die after the player has taken their turn.
func (p *player) takeTurn(die int) int {
	for i := 0; i < 3; i++ {
		p.position += die
		for p.position > 10 {
			p.position -= 10
		}
		die++
		if die == 101 {
			die = 1
		}
	}
	return die
}

func part1() {
	player1 := player{position: 7}
	player2 := player{position: 10}
	die, rolls := 1, 0
	for {
		die = player1.takeTurn(die)
		rolls += 3
		player1.score += player1.position
		if player1.score >= 1000 {
			break
		}

		die = player2.takeTurn(die)
		rolls += 3
		player2.score += player2.position
		if player2.score >= 1000 {
			break
		}
	}
	fmt.Println("Part 1 answer:", intmath.Min(player1.score, player2.score)*rolls)
}

func part2() {
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}
