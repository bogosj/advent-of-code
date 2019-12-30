package locations

import (
	"math"
	"strconv"
	"strings"

	"github.com/bogosj/advent-of-code/fileinput"
	"github.com/bogosj/advent-of-code/intmath"
)

// Locations represents the distances between locations.
type Locations struct {
	l  map[string]int
	ls map[string]bool
}

// New creates a new star chart.
func New() *Locations {
	l := Locations{}
	l.l = map[string]int{}
	l.ls = map[string]bool{}
	return &l
}

func atoiOrPanic(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func (l *Locations) allPaths() (ret [][]string) {
	var names []string
	for k := range l.ls {
		names = append(names, k)
	}
	var idxs []int
	for i := range names {
		idxs = append(idxs, i)
	}
	for _, idxPath := range intmath.Permutations(idxs) {
		var p []string
		for _, i := range idxPath {
			p = append(p, names[i])
		}
		ret = append(ret, p)
	}
	return
}

func (l *Locations) pathDistance(p []string) (ret int) {
	for i := 0; i < len(p)-1; i++ {
		ret += l.l[p[i]+"|"+p[i+1]]
	}
	return
}

// ShortestPath determines the shortest path between all locations.
func (l *Locations) ShortestPath() (path string, dist int) {
	dist = math.MaxInt32
	for _, p := range l.allPaths() {
		d := l.pathDistance(p)
		if d < dist {
			dist = d
			path = strings.Join(p, " -> ")
		}
	}
	return
}

// LongestPath determines the shortest path between all locations.
func (l *Locations) LongestPath() (path string, dist int) {
	dist = 0
	for _, p := range l.allPaths() {
		d := l.pathDistance(p)
		if d > dist {
			dist = d
			path = strings.Join(p, " -> ")
		}
	}
	return
}

// Load reads the distance chart into memory.
func (l *Locations) Load(p string) {
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		f := strings.Fields(line)
		l.ls[f[0]] = true
		l.ls[f[2]] = true
		l.l[f[0]+"|"+f[2]] = atoiOrPanic(f[4])
		l.l[f[2]+"|"+f[0]] = atoiOrPanic(f[4])
	}
	l.allPaths()
}
