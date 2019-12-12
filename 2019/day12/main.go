package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

type moon struct {
	x, y, z    int
	vX, vY, vZ int
}

func (m *moon) energy() int {
	return (abs(m.x) + abs(m.y) + abs(m.z)) * (abs(m.vX) + abs(m.vY) + abs(m.vZ))
}

func (m *moon) applyGravity(om *moon) {
	if m.x > om.x {
		m.vX--
		om.vX++
	} else if m.x < om.x {
		m.vX++
		om.vX--
	}
	if m.y > om.y {
		m.vY--
		om.vY++
	} else if m.y < om.y {
		m.vY++
		om.vY--
	}
	if m.z > om.z {
		m.vZ--
		om.vZ++
	} else if m.z < om.z {
		m.vZ++
		om.vZ--
	}
}

func (m *moon) applyVelocity() {
	m.x += m.vX
	m.y += m.vY
	m.z += m.vZ
}

func (m *moon) String() string {
	return fmt.Sprintf(
		"pos=<x=% 3d, y=% 3d, z=% 3d>, vel=<x=% 3d, y=% 3d, z=% 3d>",
		m.x, m.y, m.z, m.vX, m.vY, m.vZ)
}

func part1() {
	moons := input("input.txt")
	for range make([]bool, 1000) {
		for i := range moons {
			for j := i + 1; j < len(moons); j++ {
				moons[i].applyGravity(&moons[j])
			}
		}
		for i := range moons {
			moons[i].applyVelocity()
		}
	}

	fmt.Println("After 1000 steps:")
	energy := 0
	for _, moon := range moons {
		fmt.Println(moon.String())
		energy += moon.energy()
	}
	fmt.Println("Total energy:", energy)
}

func part2() {

}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}

func input(n string) (ret []moon) {
	fields := []rune{'<', 'x', 'y', 'z', ',', ' ', '=', '>'}
	lines := strings.Split(rawinput(n), "\n")
	for _, line := range lines {
		f := strings.FieldsFunc(line, func(r rune) bool {
			for _, fr := range fields {
				if r == fr {
					return true
				}
			}
			return false
		})
		x, _ := strconv.Atoi(f[0])
		y, _ := strconv.Atoi(f[1])
		z, _ := strconv.Atoi(f[2])
		ret = append(ret, moon{x, y, z, 0, 0, 0})
	}
	return ret
}

func rawinput(n string) string {
	data, _ := ioutil.ReadFile(n)
	return string(data)
}
