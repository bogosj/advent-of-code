package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

const (
	x = iota
	y
	z
)

type particle struct {
	p, v, a []int
	dead    bool
}

func (p *particle) String() string {
	return fmt.Sprintf("p=%v v=%v a=%v", p.p, p.v, p.a)
}

func (p *particle) locKey() string {
	return fmt.Sprintf("%v|%v|%v", p.p[x], p.p[y], p.p[z])
}

func (p *particle) distance() int {
	return intmath.Abs(p.p[x]) + intmath.Abs(p.p[y]) + intmath.Abs(p.p[z])
}

func (p *particle) advance() {
	p.v[x] += p.a[x]
	p.v[y] += p.a[y]
	p.v[z] += p.a[z]
	p.p[x] += p.v[x]
	p.p[y] += p.v[y]
	p.p[z] += p.v[z]
}

func makeVals(s string) (ret []int) {
	for _, r := range []string{"p", "v", "a", "=", "<", ">"} {
		s = strings.ReplaceAll(s, r, "")
	}
	for _, f := range strings.FieldsFunc(s, func(r rune) bool { return r == ',' }) {
		ret = append(ret, intmath.Atoi(f))
	}
	return
}

func input() (ret []*particle) {
	for _, line := range fileinput.ReadLines("input.txt") {
		p := particle{}
		f := strings.Fields(line)
		p.p = makeVals(f[0])
		p.v = makeVals(f[1])
		p.a = makeVals(f[2])
		ret = append(ret, &p)
	}
	return
}

func resolveCollisions(particles []*particle) {
	locs := map[string]int{}
	for _, p := range particles {
		if !p.dead {
			locs[p.locKey()]++
		}
	}
	for _, p := range particles {
		if locs[p.locKey()] > 1 {
			p.dead = true
		}
	}
	return
}

func aliveParticles(ps []*particle) (ret int) {
	for _, p := range ps {
		if !p.dead {
			ret++
		}
	}
	return
}

func part1() {
	particles := input()
	for i := 0; i < 500; i++ {
		for _, p := range particles {
			p.advance()
		}
	}
	minParticle := particles[0]
	for _, p := range particles {
		if p.distance() < minParticle.distance() {
			minParticle = p
		}
	}
	for i := range particles {
		if particles[i] == minParticle {
			fmt.Printf("Particle %d is the closest\n", i)
		}
	}
}

func part2() {
	particles := input()
	for i := 0; i < 1000; i++ {
		for _, p := range particles {
			p.advance()
		}
		resolveCollisions(particles)
	}
	fmt.Printf("There are %d alive particles after expansion\n", aliveParticles(particles))
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
