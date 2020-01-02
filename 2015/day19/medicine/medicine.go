package medicine

import (
	"math"
	"strings"

	"github.com/bogosj/advent-of-code/fileinput"
)

// Machine represents a medicine making machine.
type Machine struct {
	molecule      string
	replacements  map[string][]string
	inverted      bool
	iReplacements map[string][]string
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

func (m *Machine) invertReplacements() {
	m.iReplacements = map[string][]string{}
	for k, v := range m.replacements {
		for _, vv := range v {
			m.iReplacements[vv] = append(m.iReplacements[vv], k)
		}
	}
}

// New creates a new machine based on the instructions at the provided path.
func New(p string) *Machine {
	m := Machine{}
	m.load(p)
	m.invertReplacements()
	return &m
}

// Calibrate returns the number of molecules that can be created with one replacement.
func (m *Machine) Calibrate() int {
	return len(m.applyAllTransforms(buildState{s: m.molecule}))
}

type buildState struct {
	s     string
	steps int
}

func (m *Machine) applyAllTransforms(state buildState) (ret []buildState) {
	possible := map[string]int{}
	repl := m.replacements
	if m.inverted {
		repl = m.iReplacements
	}
	for k, v := range repl {
		s := strings.Split(state.s, k)
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
	}
	for k := range possible {
		ret = append(ret, buildState{s: k, steps: state.steps + 1})
	}
	return
}

func (m *Machine) decompose(state buildState) (ret []buildState) {
	if state.s == "e" {
		return []buildState{state}
	}
	newStates := m.applyAllTransforms(state)
	for _, newState := range newStates {
		decomposed := m.decompose(newState)
		ret = append(ret, decomposed...)
	}
	return
}

// Build takes the required molecule and takes all possible decomposition paths to find the shortest.
func (m *Machine) Build() int {
	m.inverted = true
	states := m.decompose(buildState{s: m.molecule})
	min := math.MaxInt32
	for _, state := range states {
		if state.steps < min {
			min = state.steps
		}
	}
	return min
}
