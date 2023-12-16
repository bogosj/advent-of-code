package wizardsim

import (
	"errors"
	"math"
)

type character struct {
	hp, damage, mana, manaUsed              int
	rechargeTicks, poisonTicks, shieldTicks int
}

type fightState struct {
	boss, fighter character
	bossTurn      bool
}

// Simulate runs all possible paths of the wizard game BFS and returns the minimum mana required to win.
func Simulate(isHard bool) int {
	b := makeBoss()
	f := makeFighter()
	minMana := math.MaxInt32
	states := []fightState{{boss: b, fighter: f}}
	for len(states) > 0 {
		state := states[0]
		states = states[1:]
		if state.fighter.manaUsed > minMana {
			continue
		}
		if state.boss.hp <= 0 {
			if minMana > state.fighter.manaUsed {
				minMana = state.fighter.manaUsed
			}
			continue
		}
		if state.fighter.hp <= 0 {
			// lost, game over, do nothing
			continue
		}
		ns := newStates(state, isHard)
		states = append(states, ns...)
	}
	return minMana
}

func (c *character) useMana(i int) error {
	if c.mana < i {
		return errors.New("not enough mana")
	}
	c.mana -= i
	c.manaUsed += i
	return nil
}

func newStates(state fightState, isHard bool) (ret []fightState) {
	if isHard && !state.bossTurn {
		state.fighter.hp--
	}
	if state.fighter.hp <= 0 {
		return
	}
	if state.boss.poisonTicks > 0 {
		state.boss.poisonTicks--
		state.boss.hp -= 3
		if state.boss.hp <= 0 {
			ret = append(ret, state)
			return
		}
	}
	if state.fighter.shieldTicks > 0 {
		state.fighter.shieldTicks--
	}
	if state.fighter.rechargeTicks > 0 {
		state.fighter.mana += 101
		state.fighter.rechargeTicks--
	}
	if state.bossTurn {
		ns := state
		ns.bossTurn = false
		if ns.fighter.shieldTicks > 0 {
			ns.fighter.hp -= ns.boss.damage - 7
		} else {
			ns.fighter.hp -= ns.boss.damage
		}
		if ns.fighter.hp > 0 {
			ret = append(ret, ns)
		}
		return
	}

	state.bossTurn = true
	var ns fightState
	if state.boss.poisonTicks == 0 {
		ns = state
		if err := ns.fighter.useMana(173); err == nil {
			ns.boss.poisonTicks = 6
			ret = append(ret, ns)
		}
	}
	if state.fighter.shieldTicks == 0 {
		ns = state
		if err := ns.fighter.useMana(113); err == nil {
			ns.fighter.shieldTicks = 6
			ret = append(ret, ns)
		}
	}
	if state.fighter.rechargeTicks == 0 {
		ns = state
		if err := ns.fighter.useMana(229); err == nil {
			ns.fighter.rechargeTicks = 5
			ret = append(ret, ns)
		}
	}
	ns = state
	if err := ns.fighter.useMana(73); err == nil {
		ns.fighter.hp += 2
		ns.boss.hp -= 2
		ret = append(ret, ns)
	}
	ns = state
	if err := ns.fighter.useMana(53); err == nil {
		ns.boss.hp -= 4
		ret = append(ret, ns)
	}
	return
}

func makeFighter() character {
	return character{hp: 50, mana: 500}
}

func makeBoss() character {
	return character{hp: 51, damage: 9}
}
