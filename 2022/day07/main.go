package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type file struct {
	name string
	size int
}

type directory struct {
	name    string
	subdirs []*directory
	parent  *directory
	files   []file
}

func (d *directory) allSubDirectories() []*directory {
	ret := []*directory{}
	for _, d := range d.subdirs {
		ret = append(ret, d)
		ret = append(ret, d.allSubDirectories()...)
	}
	return ret
}

func (d *directory) size() int {
	sum := 0
	for _, f := range d.files {
		sum += f.size
	}
	for _, d := range d.subdirs {
		sum += d.size()
	}
	return sum
}

func constructRoot(in []string) directory {
	root := &directory{name: "/"}
	cwd := root
	for _, line := range in {
		if strings.HasPrefix(line, "$ cd") {
			newDir := strings.Split(line, " ")[2]
			if newDir == ".." {
				cwd = cwd.parent
			} else {
				for _, subDir := range cwd.subdirs {
					if subDir.name == newDir {
						cwd = subDir
						break
					}
				}
			}
		} else if strings.HasPrefix(line, "$ ls") {
			//noop
		} else if strings.HasPrefix(line, "dir") {
			newDir := &directory{}
			newDir.name = strings.Split(line, " ")[1]
			newDir.parent = cwd
			cwd.subdirs = append(cwd.subdirs, newDir)
		} else {
			newFile := file{
				size: intmath.Atoi(strings.Split(line, " ")[0]),
				name: strings.Split(line, " ")[1],
			}
			cwd.files = append(cwd.files, newFile)
		}
	}
	return *root
}

func part1(in []string) {
	root := constructRoot(in)
	sum := 0
	for _, d := range root.allSubDirectories() {
		if d.size() <= 100000 {
			sum += d.size()
		}
	}
	fmt.Printf("The sum of the size is: %d\n", sum)
}

func part2(in []string) {
	root := constructRoot(in)
	freeSpace := 70000000 - root.size()
	neededSpace := 30000000 - freeSpace
	allSubs := root.allSubDirectories()
	sort.Slice(allSubs, func(i, j int) bool {
		return allSubs[i].size() < allSubs[j].size()
	})
	for _, d := range allSubs {
		if d.size() >= neededSpace {
			fmt.Printf("The smallest directory you can delete is: %d\n", d.size())
			return
		}
	}
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
