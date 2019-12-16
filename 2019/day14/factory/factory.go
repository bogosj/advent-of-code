package factory

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bogosj/advent-of-code/2019/fileinput"
)

// Factory represents a factory that turns ore into fuel.
type Factory struct {
	inventory map[string]int
	reactions map[string]reaction
}

func (f *Factory) String() (ret string) {
	for k, v := range f.reactions {
		ret += fmt.Sprintf("%v: %v\n", k, v)
	}
	return
}

// New creates a new factory from the path of a provided input file.
func New(in string) *Factory {
	f := Factory{}
	f.inventory = map[string]int{}
	f.reactions = map[string]reaction{}
	for _, r := range input(in) {
		n := r.Output.Name
		f.reactions[n] = r
	}
	return &f
}

// Ore determines how much ore is needed for a provided amount of fuel.
func (f *Factory) Ore(fuel int) int {
	i := map[string]int{}
	i["FUEL"] = fuel

make:
	for {
		for m := range i {
			if m != "ORE" && i[m] > 0 {
				a := (i[m]-1)/f.reactions[m].Output.Quantity + 1
				i[m] -= f.reactions[m].Output.Quantity * a

				for _, c := range f.reactions[m].Input {
					i[c.Name] += c.Quantity * a
				}
				continue make
			}
		}
		return i["ORE"]
	}
}

type chemical struct {
	Quantity int
	Name     string
}

type reaction struct {
	Input  []chemical
	Output chemical
}

func input(n string) (ret []reaction) {
	lines := fileinput.ReadLines("input.txt")
	for _, line := range lines {
		f := strings.FieldsFunc(line, func(r rune) bool {
			return r == ',' || r == ' ' || r == '=' || r == '>'
		})
		var c []chemical
		for len(f) > 0 {
			q, err := strconv.Atoi(f[0])
			if err != nil {
				fmt.Println(err)
			}
			c = append(c, chemical{q, f[1]})
			f = f[2:]
		}
		ret = append(ret, reaction{c[:len(c)-1], c[len(c)-1]})
	}
	return ret
}
