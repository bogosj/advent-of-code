package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/2019/day22/deck"
	"github.com/bogosj/advent-of-code/fileinput"
)

const (
	smarterPersonPart2Solution = "https://raw.githubusercontent.com/twattanawaroon/adventofcode/0735e64bd9c6f4b3d014075216b3412b39cc9ec7/2019/q22b.py"
	quote                      = `
For part 2... I'm not smart enough to figure this out, we're going to:
- find the solution of someone who is smarter than me and understands this math
- run curl to download it (pinned to a specific commit)
- run it
- print the answer

Their code can be found at %s

`
)

func part1() {
	d := deck.New(10007)
	d.RunInstructions("input.txt")
	fmt.Printf("Card at position %d is %d\n", 2019, d.PosOfCard(2019))
}

func makeTempDir() string {
	out, err := exec.Command("mktemp", "-d").CombinedOutput()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

func getSolution(tempDir string) string {
	path := fmt.Sprintf("%s/part2.py", tempDir)
	out, err := exec.Command("curl", smarterPersonPart2Solution).Output()
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(path, out, 0644)
	if err != nil {
		panic(err)
	}
	return path
}

func part2() {
	fmt.Printf(quote, smarterPersonPart2Solution)
	solutionPath := getSolution(makeTempDir())

	cmd := exec.Command("/usr/bin/python", solutionPath)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, strings.Join(fileinput.ReadLines("input.txt"), "\n"))
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in", time.Since(start))
}
