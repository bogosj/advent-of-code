package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

// 21760 too lows

const (
	immune = iota
	infection
)

type group struct {
	damage, hp        int
	units, initiative int
	team              int
	attackType        string
	immune, weak      []string
}

func (g *group) String() string {
	return fmt.Sprintf("i%d w/ %d units w/ %d HP", g.initiative, g.units, g.hp)
}

func (g *group) effectivePower() int {
	return g.units * g.damage
}

func (g *group) damageCausedBy(attacker *group) (ret int) {
	ret = attacker.effectivePower()
	for _, t := range g.immune {
		if attacker.attackType == t {
			ret = 0
		}
	}
	for _, t := range g.weak {
		if attacker.attackType == t {
			ret *= 2
		}
	}
	return
}

func (g *group) attack(og *group) {
	if og == nil {
		return
	}
	unitsKilled := intmath.Min(og.damageCausedBy(g)/og.hp, og.units)
	og.units -= unitsKilled
}

func (g *group) targets(groups []*group) (ret []*group) {
	for _, og := range groups {
		if og.team == g.team {
			continue
		}
		ret = append(ret, og)
	}
	sort.Slice(ret, func(i, j int) bool {
		if ret[i].damageCausedBy(g) == ret[j].damageCausedBy(g) {
			if ret[i].effectivePower() == ret[j].effectivePower() {
				return ret[i].initiative > ret[j].initiative
			}
			return ret[i].effectivePower() > ret[j].effectivePower()
		}
		return ret[i].damageCausedBy(g) > ret[j].damageCausedBy(g)
	})
	return ret
}

func input(s string) (ret []*group) {
	atoi := intmath.Atoi
	team := immune
	lines := fileinput.ReadLines(s)
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			i++
			team = infection
			continue
		}
		ng := group{}
		f := strings.Fields(line)
		ng.units = atoi(f[0])
		f = f[4:]
		ng.hp = atoi(f[0])
		f = f[3:]
		for f[0] != "with" {
			if f[0][0] == 'w' || f[0][1] == 'w' {
				f = f[2:]
				for {
					s := f[0]
					s = strings.ReplaceAll(s, ",", "")
					s = strings.ReplaceAll(s, ";", "")
					s = strings.ReplaceAll(s, ")", "")
					ng.weak = append(ng.weak, s)
					f = f[1:]
					if f[0] == "immune" || f[0] == "with" {
						break
					}
				}
			} else {
				f = f[2:]
				for {
					s := f[0]
					s = strings.ReplaceAll(s, ",", "")
					s = strings.ReplaceAll(s, ";", "")
					s = strings.ReplaceAll(s, ")", "")
					ng.immune = append(ng.immune, s)
					f = f[1:]
					if f[0] == "weak" || f[0] == "with" {
						break
					}
				}
			}
		}
		ng.damage = atoi(f[5])
		ng.attackType = f[6]
		ng.initiative = atoi(f[10])
		ng.team = team
		ret = append(ret, &ng)
	}
	return
}

func allOfOneTeam(groups []*group) bool {
	firstTeam := groups[0].team
	for i := 1; i < len(groups); i++ {
		if groups[i].team != firstTeam {
			return false
		}
	}
	return true
}

func removeDeadGroups(groups []*group) (ret []*group) {
	for _, g := range groups {
		if g.units > 0 {
			ret = append(ret, g)
		}
	}
	return
}

func fightOneRound(groups []*group) []*group {
	// sort by power and initiative
	sort.Slice(groups, func(i, j int) bool {
		if groups[i].effectivePower() == groups[j].effectivePower() {
			return groups[i].initiative > groups[j].initiative
		}
		return groups[i].effectivePower() > groups[j].effectivePower()
	})
	// select targets
	targets := map[*group]*group{}
	targeted := map[*group]bool{}
	for _, g := range groups {
		for _, t := range g.targets(groups) {
			if !targeted[t] {
				if t.damageCausedBy(g) == 0 {
					continue
				}
				targeted[t] = true
				targets[g] = t
				break
			}
		}
	}
	// sort by decreasing initiative
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].initiative > groups[j].initiative
	})
	// perform attack
	for _, g := range groups {
		if g.units == 0 {
			continue
		}
		g.attack(targets[g])
	}
	return removeDeadGroups(groups)
}

func performCombat(groups []*group) (winners []*group) {
	var rounds int
	for rounds < 10000 {
		if allOfOneTeam(groups) {
			for _, g := range groups {
				winners = append(winners, g)
			}
			return
		}
		groups = fightOneRound(groups)
		rounds++
	}
	return nil
}

func part1() {
	groups := input("input.txt")
	winners := performCombat(groups)
	var total int
	for _, g := range winners {
		total += g.units
	}
	fmt.Printf("The winning army has %d units.\n", total)
}

func boost(groups []*group, amt int) {
	for _, g := range groups {
		if g.team == immune {
			g.damage += amt
		}
	}
}

func part2() {
	i := sort.Search(1_000_000, func(n int) bool {
		groups := input("input.txt")
		boost(groups, n)
		winners := performCombat(groups)
		if winners == nil {
			return false
		}
		return winners[0].team == immune
	})
	fmt.Printf("A boost of %d is the minimum required for the immune system to win.\n", i)

	groups := input("input.txt")
	boost(groups, i)
	winners := performCombat(groups)
	var total int
	for _, g := range winners {
		total += g.units
	}
	fmt.Printf("The winning army has %d units.\n", total)
}

func main() {
	start := time.Now()
	part1()
	fmt.Println("Part 1 done in:", time.Since(start))
	start = time.Now()
	part2()
	fmt.Println("Part 2 done in:", time.Since(start))
}
