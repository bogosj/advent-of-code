package rpg

import (
	"github.com/bogosj/advent-of-code/intmath"
)

type character struct {
	hp, damage, armor, cost int
	name                    string
}

func (c *character) String() string {
	return c.name
}

func (c *character) attack(o *character) {
	d := intmath.Max(1, c.damage-o.armor)
	o.hp -= d
}

// CostOfCheapestWinningFighter simulates all possible battles and returns the
// lowest cost of the equipment required to win.
func CostOfCheapestWinningFighter() int {
	var costs []int
	for _, f := range allWinningFighters("Me") {
		costs = append(costs, f.cost)
	}
	return intmath.Min(costs...)
}

func allWinningFighters(n string) (ret []*character) {
	for c := range allItemCombinations() {
		b := makeBoss()
		f := makeFighter(c)
		winner := f.fight(b)
		if winner.name == n {
			ret = append(ret, f)
		}
	}
	return
}

func (c *character) fight(o *character) *character {
	for {
		c.attack(o)
		if o.hp <= 0 {
			return c
		}
		o.attack(c)
		if c.hp <= 0 {
			return o
		}
	}
}

func makeFighter(items []item) *character {
	c := character{hp: 100, name: "Me"}
	for _, i := range items {
		c.damage += i.damage
		c.armor += i.armor
		c.cost += i.cost
	}
	return &c
}

func makeBoss() *character {
	c := character{name: "Boss", hp: 104, damage: 8, armor: 1}
	return &c
}
