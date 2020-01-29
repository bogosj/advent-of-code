package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

type nanobot struct {
	x, y, z, r int
}

func (n *nanobot) String() string {
	return fmt.Sprintf("<%d %d %d> | %d", n.x, n.y, n.z, n.r)
}

func (n *nanobot) canSee(on *nanobot) bool {
	dist := intmath.Abs(n.x-on.x) + intmath.Abs(n.y-on.y) + intmath.Abs(n.z-on.z)
	return n.r >= dist
}

func (n *nanobot) distToOrigin() int {
	return intmath.Abs(n.x) + intmath.Abs(n.y) + intmath.Abs(n.z)
}

func input() (ret []*nanobot) {
	for _, line := range fileinput.ReadLines("input.txt") {
		b := &nanobot{}
		f := strings.Split(line, " ")
		b.r = intmath.Atoi(strings.Split(f[1], "=")[1])
		p := strings.Split(f[0], "=")[1]
		p = strings.ReplaceAll(p, "<", "")
		p = strings.ReplaceAll(p, ">", "")
		pos := strings.Split(p, ",")
		b.x, b.y, b.z = intmath.Atoi(pos[0]), intmath.Atoi(pos[1]), intmath.Atoi(pos[2])
		ret = append(ret, b)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].r > ret[j].r
	})
	return
}

func bestPoint(bots []*nanobot) nanobot {
	max, min, ranges := minMax(bots)
	bestLocation := nanobot{}
	grain := int(math.Pow(2, 23))
	for {
		for x := min.x; x < max.x; x += grain {
			for y := min.y; y < max.y; y += grain {
				for z := min.z; z < max.z; z += grain {
					currPoint := nanobot{x: x, y: y, z: z}
					for _, bot := range bots {
						if bot.canSee(&currPoint) {
							currPoint.r++
						}
					}
					if currPoint.r == bestLocation.r {
						if currPoint.distToOrigin() < bestLocation.distToOrigin() {
							bestLocation = currPoint
						}
					}
					if currPoint.r > bestLocation.r {
						bestLocation = currPoint
					}
				}
			}
		}
		fmt.Printf("grain:%v min:%v max:%v ranges:%v best:%v\n", grain, min, max, ranges, bestLocation)
		if grain == 1 {
			break
		}
		grain /= 2
		ranges.x /= 2
		ranges.y /= 2
		ranges.z /= 2
		min.x = bestLocation.x - ranges.x/2
		min.y = bestLocation.y - ranges.y/2
		min.z = bestLocation.z - ranges.z/2
		max.x = bestLocation.x + ranges.x/2
		max.y = bestLocation.y + ranges.y/2
		max.z = bestLocation.z + ranges.z/2
	}
	fmt.Printf("%v\n", bestLocation)
	fmt.Printf("Distance to origin for best point is: %d\n", bestLocation.distToOrigin())
	return bestLocation
}

func minMax(bots []*nanobot) (nanobot, nanobot, nanobot) {
	var min, max, ranges nanobot
	for _, bot := range bots {
		if bot.x < min.x {
			min.x = bot.x
		}
		if bot.y < min.y {
			min.y = bot.y
		}
		if bot.z < min.z {
			min.z = bot.z
		}
		if bot.x > max.x {
			max.x = bot.x
		}
		if bot.y > max.y {
			max.y = bot.y
		}
		if bot.z > max.z {
			max.z = bot.z
		}
	}
	ranges.x = max.x - min.x
	ranges.y = max.y - min.y
	ranges.z = max.z - min.z
	return max, min, ranges
}

func part1() {
	bots := input()
	var c int
	for _, bot := range bots {
		if bots[0].canSee(bot) {
			c++
		}
	}
	fmt.Printf("There are %d bots in range of the bot with the biggest range.\n", c)
}

func part2() {
	bestPoint(input())
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
