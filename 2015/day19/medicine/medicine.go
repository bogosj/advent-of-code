package medicine

import (
	"fmt"
	"strings"

	"github.com/bogosj/advent-of-code/fileinput"
)

// Machine represents a medicine making machine.
type Machine struct {
	molecule     string
	replacements map[string][]string
}

func (m *Machine) load(p string) {
	m.replacements = map[string][]string{}
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		f := strings.FieldsFunc(line, func(r rune) bool {
			return r == '=' || r == '>'
		})
		if len(f) == 0 {
			continue
		}
		f0 := strings.TrimSpace(f[0])
		if len(f) == 1 {
			m.molecule = f0
		}
		if len(f) == 2 {
			f1 := strings.TrimSpace(f[1])
			m.replacements[f0] = append(m.replacements[f0], f1)
		}
	}
}

// New creates a new machine based on the instructions at the provided path.
func New(p string) *Machine {
	m := Machine{}
	m.load(p)
	return &m
}

// Calibrate returns the number of molecules that can be created with one replacement.
func (m *Machine) Calibrate() int {
	possible := map[string]int{}
	c := 0
	for k, v := range m.replacements {
		s := strings.Split(m.molecule, k)
		if len(s) == 1 {
			continue
		}
		for i := 0; i < len(s)-1; i++ {
			for _, replacement := range v {
				newStr := ""
				for j, ss := range s {
					newStr += ss
					if j == i {
						newStr += replacement
					} else if j < len(s)-1 {
						newStr += k
					}
				}
				possible[newStr]++
			}
		}
		c++
		if c > 1000000 {
			return -1
		}

	}
	for k, v := range possible {
		if v > 1 {
			fmt.Println(v, k)
		}
	}
	return len(possible)
}
