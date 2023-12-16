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

// CostOfCheapestWinner simulates all possible battles and returns the
// lowest cost of the equipment required to win.
func CostOfCheapestWinner() int {
	var costs []int
	for _, f := range simulate(true) {
		costs = append(costs, f.cost)
	}
	return intmath.Min(costs...)
}

// CostOfMostExpensiveLoser simulates all possible battles and returns the
// highest cost of the equipment where you still lowse.
func CostOfMostExpensiveLoser() int {
	var costs []int
	for _, f := range simulate(false) {
		costs = append(costs, f.cost)
	}
	return intmath.Max(costs...)
}

func simulate(returnWinners bool) (ret []*character) {
	for c := range allItemCombinations() {
		b := makeBoss()
		f := makeFighter(c)
		winner := f.fight(b)
		if returnWinners {
			if winner == f {
				ret = append(ret, winner)
			}
		} else {
			if winner == b {
				ret = append(ret, f)
			}
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
