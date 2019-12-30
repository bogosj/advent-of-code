package wires

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bogosj/advent-of-code/fileinput"
)

const (
	opAnd = iota
	opNot
	opOr
	opLShift
	opRShift
	opSet
)

type operation struct {
	src1, src2, dst string
	op              uint16
}

// Wires represents a toy for Bobby Tables.
type Wires struct {
	w   map[string]uint16
	ops []operation
	in  map[string]bool
}

func (w *Wires) String() (ret string) {
	for k, v := range w.w {
		ret += fmt.Sprintf("%v: %d\n", k, v)
	}
	return
}

// ValueOf returns the value of a given wire.
func (w *Wires) ValueOf(s string) uint16 {
	return w.w[s]
}

// New creates a new Wires instance.
func New() *Wires {
	w := Wires{}
	w.w = map[string]uint16{}
	w.in = map[string]bool{"": true}
	return &w
}

func atoiOrPanic(s string) uint16 {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return uint16(i)
}

func (w *Wires) ident(s string) {
	i, err := strconv.Atoi(s)
	if err == nil {
		w.w[s] = uint16(i)
		w.in[s] = true
	}
}

func (w *Wires) processLine(s string) (ret operation) {
	f := strings.Fields(s)
	ret.dst = f[len(f)-1]

	if len(f) == 3 {
		ret.op = opSet
		ret.src1 = f[0]
		w.ident(ret.src1)
		return
	}

	if f[0][0] == 'N' {
		ret.op = opNot
		ret.src1 = f[1]
		w.ident(ret.src1)
		return
	}

	ret.src1 = f[0]
	ret.src2 = f[2]
	switch f[1][0] {
	case 'A':
		ret.op = opAnd
	case 'O':
		ret.op = opOr
	case 'L':
		ret.op = opLShift
	case 'R':
		ret.op = opRShift
	}
	w.ident(ret.src1)
	w.ident(ret.src2)
	return
}

func (w *Wires) getVal(s string) uint16 {
	i, err := strconv.Atoi(s)
	if err != nil {
		return w.w[s]
	}
	return uint16(i)
}

// Apply takes the loaded instructions and applies them to the wires.
func (w *Wires) Apply() {
	for len(w.ops) > 0 {
		op := w.ops[0]
		w.ops = w.ops[1:]
		if w.in[op.src1] && w.in[op.src2] {
			switch op.op {
			case opSet:
				w.w[op.dst] = w.w[op.src1]
			case opAnd:
				w.w[op.dst] = w.w[op.src1] & w.w[op.src2]
			case opOr:
				w.w[op.dst] = w.w[op.src1] | w.w[op.src2]
			case opNot:
				w.w[op.dst] = ^w.w[op.src1]
			case opLShift:
				w.w[op.dst] = w.w[op.src1] << w.w[op.src2]
			case opRShift:
				w.w[op.dst] = w.w[op.src1] >> w.w[op.src2]
			}
			w.in[op.dst] = true
		} else {
			w.ops = append(w.ops, op)
		}
	}
}

// Load reads the instructions for setting up the toy.
func (w *Wires) Load(p string) {
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		w.ops = append(w.ops, w.processLine(line))
	}
}
