package lights

import (
	"strconv"
	"strings"

	"github.com/bogosj/advent-of-code/fileinput"
)

// Lights represents an array of lights.
type Lights struct {
	l [][]bool
}

// New returns a new array of lights.
func New() *Lights {
	l := Lights{}
	for i := 0; i < 1000; i++ {
		r := make([]bool, 1000)
		l.l = append(l.l, r)
	}
	return &l
}

// LitLights returns the number of lights in the array that are lit.
func (l *Lights) LitLights() (ret int) {
	for _, row := range l.l {
		for _, cell := range row {
			if cell {
				ret++
			}
		}
	}
	return
}

func pair(s string) (ret []int) {
	for _, ss := range strings.Split(s, ",") {
		i, err := strconv.Atoi(ss)
		if err != nil {
			panic(err)
		}
		ret = append(ret, i)
	}
	return
}

func (l *Lights) runCommand(cmd, pair1, pair2 string) {
	p1 := pair(pair1)
	p2 := pair(pair2)
	for y := p1[0]; y <= p2[0]; y++ {
		for x := p1[1]; x <= p2[1]; x++ {
			switch cmd {
			case "toggle":
				l.l[y][x] = !l.l[y][x]
			case "on":
				l.l[y][x] = true
			case "off":
				l.l[y][x] = false
			}
		}
	}
}

// RunInstructions runs the instructions in the provided file.
func (l *Lights) RunInstructions(s string) {
	lines := fileinput.ReadLines(s)
	for _, line := range lines {
		f := strings.Fields(line)
		if f[0] == "toggle" {
			l.runCommand(f[0], f[1], f[3])
		} else {
			l.runCommand(f[1], f[2], f[4])
		}
	}
}
