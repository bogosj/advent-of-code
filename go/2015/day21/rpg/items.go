package rpg

import (
	combinations "github.com/mxschmitt/golang-combinations"
)

/*
Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3
*/

type item struct {
	cost, damage, armor int
	name                string
}

var (
	weapons = map[string]item{
		"Dagger":     {name: "Dagger", cost: 8, damage: 4},
		"Shortsword": {name: "Shortsword", cost: 10, damage: 5},
		"Warhammer":  {name: "Warhammer", cost: 25, damage: 6},
		"Longsword":  {name: "Longsword", cost: 40, damage: 7},
		"Greataxe":   {name: "Greataxe", cost: 74, damage: 8},
	}
	armor = map[string]item{
		"Leather":    {name: "Leather", cost: 13, armor: 1},
		"Chainmail":  {name: "Chainmail", cost: 31, armor: 2},
		"Splintmail": {name: "Splintmail", cost: 53, armor: 3},
		"Bandedmail": {name: "Bandedmail", cost: 75, armor: 4},
		"Platemail":  {name: "Platemail", cost: 102, armor: 5},
	}
	rings = map[string]item{
		"Damage +1":  {name: "Damage +1", cost: 25, damage: 1},
		"Damage +2":  {name: "Damage +2", cost: 50, damage: 2},
		"Damage +3":  {name: "Damage +3", cost: 100, damage: 3},
		"Defense +1": {name: "Defense +1", cost: 20, armor: 1},
		"Defense +2": {name: "Defense +2", cost: 40, armor: 2},
		"Defense +3": {name: "Defense +3", cost: 80, armor: 3},
	}
)

func allArmorCombinations() <-chan []item {
	var n []string
	for k := range armor {
		n = append(n, k)
	}
	out := make(chan []item)
	go func(n []string) {
		defer close(out)
		out <- nil
		for _, c := range combinations.All(n) {
			if len(c) > 1 {
				continue
			}
			var i []item
			for _, s := range c {
				i = append(i, armor[s])
			}
			out <- i
		}
	}(n)
	return out
}

func allRingCombinations() <-chan []item {
	var n []string
	for k := range rings {
		n = append(n, k)
	}
	out := make(chan []item)
	go func(n []string) {
		defer close(out)
		out <- nil
		for _, c := range combinations.All(n) {
			if len(c) > 2 {
				continue
			}
			var i []item
			for _, s := range c {
				i = append(i, rings[s])
			}
			out <- i
		}
	}(n)
	return out
}

func allItemCombinations() <-chan []item {
	out := make(chan []item)
	go func() {
		defer close(out)

		for _, w := range weapons {
			for a := range allArmorCombinations() {
				for rs := range allRingCombinations() {
					ret := []item{w}
					ret = append(ret, a...)
					ret = append(ret, rs...)
					out <- ret
				}
			}
		}
	}()
	return out
}
