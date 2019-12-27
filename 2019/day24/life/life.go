package life

import (
	"github.com/bogosj/advent-of-code/2019/fileinput"
	"github.com/bogosj/advent-of-code/2019/intmath"
)

const (
	bug   = '#'
	space = '.'
)

// Life represents life on Eris
type Life struct {
	board      [][]rune
	prevBoards map[string]bool
}

// New returns a new Life instance based on the provided input file.
func New(p string) *Life {
	l := Life{}
	l.prevBoards = map[string]bool{}
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		var row []rune
		for _, r := range line {
			row = append(row, r)
		}
		l.board = append(l.board, row)
	}
	return &l
}

func (l *Life) String() (ret string) {
	for _, row := range l.board {
		for _, r := range row {
			ret += string(r)
		}
		ret += "\n"
	}
	return
}

func boardToString(b [][]rune) (ret string) {
	for _, row := range b {
		for _, r := range row {
			ret += string(r)
		}
	}
	return
}

func dupeBoard(b [][]rune) (ret [][]rune) {
	for _, row := range b {
		var nr []rune
		for _, r := range row {
			nr = append(nr, r)
		}
		ret = append(ret, nr)
	}
	return
}

func score(b [][]rune) (ret int) {
	s := boardToString(b)
	for i, c := range s {
		if c == bug {
			ret += 1 << i
		}
	}
	return
}

// Run lets life progress then returns the biodiversity score of the first repeated board.
func (l *Life) Run() int {
	for !l.advance() {
	}
	return score(l.board)
}

func (l *Life) advance() bool {
	s := boardToString(l.board)
	if _, ok := l.prevBoards[s]; ok {
		return true
	}
	l.prevBoards[s] = true

	nb := dupeBoard(l.board)
	for y := 0; y < len(nb); y++ {
		for x := 0; x < len(nb[y]); x++ {
			p := intmath.Point{X: x, Y: y}
			bugs := 0
			for _, n := range p.Neighbors() {
				if n.X >= 0 && n.Y >= 0 && n.X < len(nb[y]) && n.Y < len(nb) {
					if l.board[n.Y][n.X] == bug {
						bugs++
					}
				}
			}
			nb[y][x] = space
			if l.board[y][x] == space && (bugs == 1 || bugs == 2) {
				nb[y][x] = bug
			} else if l.board[y][x] == bug && bugs == 1 {
				nb[y][x] = bug
			}
		}
	}
	l.board = nb
	return false
}
