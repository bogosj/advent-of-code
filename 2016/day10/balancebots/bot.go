package balancebots

import (
	"sort"
	"strings"

	"github.com/bogosj/advent-of-code/intmath"
)

// Factory represents a bot factory.
type Factory struct {
	bots    map[string]*bot
	output  map[string]*bot
	LookFor []int
}

// New creates a new instance of a bot factory.
func New() *Factory {
	f := Factory{}
	f.bots = map[string]*bot{}
	f.output = map[string]*bot{}
	return &f
}

func (f *Factory) getBot(s string) *bot {
	b := f.bots[s]
	if b == nil {
		nb := bot{}
		f.bots[s] = &nb
		return &nb
	}
	return b
}

func (f *Factory) getOutput(s string) *bot {
	b := f.output[s]
	if b == nil {
		nb := bot{}
		f.output[s] = &nb
		return &nb
	}
	return b
}

// RunInstructions executes the instructions provided.
func (f *Factory) RunInstructions(inst []string) string {
	for i := 0; ; i++ {
		for _, line := range inst {
			fields := strings.Fields(line)
			if fields[0] == "value" {
				if i == 0 {
					v := intmath.Atoi(fields[1])
					b := f.getBot(fields[5])
					b.giveChip(v)
				}
				continue
			}
			giver := f.getBot(fields[1])
			if len(giver.chips) != 2 {
				continue
			}
			if f.LookFor != nil {
				if giver.chips[0] == f.LookFor[0] && giver.chips[1] == f.LookFor[1] {
					return fields[1]
				}
			}
			var r1, r2 *bot
			if fields[5] == "bot" {
				r1 = f.getBot(fields[6])
			} else {
				r1 = f.getOutput(fields[6])
			}
			if fields[10] == "bot" {
				r2 = f.getBot(fields[11])
			} else {
				r2 = f.getOutput(fields[11])
			}
			giver.giveTo(r1, r2)
		}
	}
}

type bot struct {
	chips []int
}

func (b *bot) giveChip(i int) {
	b.chips = append(b.chips, i)
	sort.Ints(b.chips)
}

func (b *bot) giveTo(ob1, ob2 *bot) {
	ob1.giveChip(b.chips[0])
	ob2.giveChip(b.chips[1])
	b.chips = nil
}
