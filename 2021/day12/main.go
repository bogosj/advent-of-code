package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
)

func makeCaves(in []string) map[string][]string {
	ret := map[string][]string{}
	for _, rooms := range in {
		pair := strings.Split(rooms, "-")
		ret[pair[0]] = append(ret[pair[0]], pair[1])
		ret[pair[1]] = append(ret[pair[1]], pair[0])
	}
	return ret
}

func roomInPath(room string, path []string) bool {
	for _, r := range path {
		if room == r {
			return true
		}
	}
	return false
}

func findPaths(caves map[string][]string) [][]string {
	re := regexp.MustCompile(`[a-z]`)
	paths := [][]string{}
	potentialPaths := [][]string{}

	potentialPaths = append(potentialPaths, []string{"start"})
	for len(potentialPaths) > 0 {
		path := potentialPaths[0]
		potentialPaths = potentialPaths[1:]
		lastRoom := path[len(path)-1]
		if lastRoom == "end" {
			paths = append(paths, path)
			continue
		}
		for _, nextRoom := range caves[lastRoom] {
			if (re.MatchString(nextRoom) && !roomInPath(nextRoom, path)) || !re.MatchString(nextRoom) || nextRoom == "end" {
				newPath := []string{}
				newPath = append(newPath, path...)
				newPath = append(newPath, nextRoom)
				potentialPaths = append(potentialPaths, newPath)
			}
		}
	}

	return paths
}

func part1(in []string) {
	caves := makeCaves(in)
	fmt.Println("Part 1 answer:", len(findPaths(caves)))
}

func part2(in []string) {
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

func input() []string {
	return fileinput.ReadLines("input.txt")
}
