package deck

import (
	"strconv"
	"strings"

	"github.com/bogosj/advent-of-code/fileinput"
)

// Deck represents a deck of space cards.
type Deck struct {
	d []int
}

// New creates a new deck of space cards.
func New(s int) *Deck {
	d := Deck{}
	d.d = make([]int, s)
	for i := 0; i < s; i++ {
		d.d[i] = i
	}
	return &d
}

// NewStack deals the deck into a new stack in reverse order.
func (d *Deck) NewStack() {
	for i := len(d.d)/2 - 1; i >= 0; i-- {
		opp := len(d.d) - 1 - i
		d.d[i], d.d[opp] = d.d[opp], d.d[i]
	}
}

// Cut cuts the deck at the provided point.
func (d *Deck) Cut(n int) {
	var nd []int
	if n > 0 {
		nd = d.d[n:]
		nd = append(nd, d.d[:n]...)
	} else {
		nd = d.d[len(d.d)+n:]
		nd = append(nd, d.d[:len(d.d)+n]...)
	}
	d.d = nd
}

// Deal deals the cards to n positions.
func (d *Deck) Deal(n int) {
	nd := make([]int, len(d.d))
	for i := 0; i < len(d.d); i++ {
		nd[(i*n)%len(d.d)] = d.d[i]
	}
	d.d = nd
}

// PosOfCard returns the position of the card of the provided value.
func (d *Deck) PosOfCard(n int) int {
	for i := 0; i < len(d.d); i++ {
		if d.d[i] == n {
			return i
		}
	}
	return -1
}

// RunInstructions runs the instructions provided at the given path.
func (d *Deck) RunInstructions(p string) {
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		inst := strings.Fields(line)
		if inst[0] == "cut" {
			n, err := strconv.Atoi(inst[1])
			if err != nil {
				panic(err)
			}
			d.Cut(n)
		} else if inst[0] == "deal" && inst[1] == "into" {
			d.NewStack()
		} else if inst[0] == "deal" && inst[1] == "with" {
			n, err := strconv.Atoi(inst[3])
			if err != nil {
				panic(err)
			}
			d.Deal(n)
		}
	}
}
