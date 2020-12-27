package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/bogosj/advent-of-code/2019/day22/deck"
)

const (
	smarterPersonPart2Solution = "https://raw.githubusercontent.com/twattanawaroon/adventofcode/0735e64bd9c6f4b3d014075216b3412b39cc9ec7/2019/q22b.py"
	quote                      = `
For part 2... I'm not smart enough to figure this out, we're going to
find the solution of someone who is smarter than me and understands this
math and run their code in a docker container.

Their code can be found at %s
`
)

func part1() {
	d := deck.New(10007)
	d.RunInstructions("input.txt")
	fmt.Printf("Card at position %d is %d\n", 2019, d.PosOfCard(2019))
}

func part2() {
	fmt.Printf(quote, smarterPersonPart2Solution)
	out, err := exec.Command("./solve_part2.sh").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nPart 2 solution: %s\n", string(out))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}
