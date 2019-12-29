package vault

import (
	"fmt"
	"math"
	"unicode"

	"github.com/bogosj/advent-of-code/2019/intmath"
	"github.com/bogosj/advent-of-code/fileinput"
)

const (
	wall  = '#'
	space = '.'
)

// Vault represents a vault with keys and doors.
type Vault struct {
	m            [][]rune
	keys         map[rune]intmath.Point
	keysToUnlock map[rune]map[rune]intmath.Bitmap
	distances    map[rune]map[rune]int
	score        map[string]int
	origin       intmath.Point
}

func isKey(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func isDoor(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func (v *Vault) copyMap() (ret [][]rune) {
	for _, r := range v.m {
		nr := []rune{}
		for _, c := range r {
			nr = append(nr, c)
		}
		ret = append(ret, nr)
	}
	return
}

type distPoint struct {
	p     intmath.Point
	dist  int
	doors intmath.Bitmap
}

func (v *Vault) findPathsForFrom(c rune, p intmath.Point, m [][]rune, doors intmath.Bitmap) {
	var dps []distPoint
	for _, p := range p.Neighbors() {
		dps = append(dps, distPoint{p: p, dist: 1})
	}

	for len(dps) > 0 {
		dp := dps[0]
		dps = dps[1:]
		n := dp.p
		doors := dp.doors
		if n.Y < 0 || n.X < 0 {
			continue
		}
		s := m[n.Y][n.X]
		if s != wall && s != c {
			if isKey(s) {
				d, ok := v.distances[c][s]
				if !ok || d > dp.dist {
					v.distances[c][s] = dp.dist
					v.keysToUnlock[c][s] = doors
				}
			}
			if isDoor(s) {
				doors = doors.Set(int(unicode.ToLower(s) - 'a'))
			}
			for _, p := range n.Neighbors() {
				dps = append(dps, distPoint{p: p, dist: dp.dist + 1, doors: doors})
			}
			m[n.Y][n.X] = wall
		}
	}
}

func (v *Vault) findPathsFor(c rune) {
	m := v.copyMap()
	if c == '@' {
		v.findPathsForFrom(c, v.origin, m, intmath.Bitmap{})
	} else {
		v.findPathsForFrom(c, v.keys[c], m, intmath.Bitmap{})
	}
}

func (v *Vault) findPaths() {
	v.distances = map[rune]map[rune]int{}
	v.keysToUnlock = map[rune]map[rune]intmath.Bitmap{}
	for c := 'a'; c <= 'z'; c++ {
		if _, ok := v.keys[c]; ok {
			v.distances[c] = map[rune]int{}
			v.keysToUnlock[c] = map[rune]intmath.Bitmap{}
			v.findPathsFor(c)
		}
	}
	v.distances['@'] = map[rune]int{}
	v.keysToUnlock['@'] = map[rune]intmath.Bitmap{}
	v.findPathsFor('@')
}

func (v *Vault) findKeyPoints() {
	v.keys = map[rune]intmath.Point{}
	for y := range v.m {
		for x := range v.m[y] {
			p := intmath.Point{X: x, Y: y}
			c := v.m[y][x]
			if c >= 'a' && c <= 'z' {
				v.keys[c] = p
			}
			if c == '@' {
				v.origin = p
			}
		}
	}
}

func toLetters(b intmath.Bitmap) (ret string) {
	for _, i := range b.OneBits() {
		ret += string('a' + i)
	}
	return
}

func (v *Vault) visibleKeys(fromKey rune, keysNeeded, openDoors intmath.Bitmap) (ret []rune) {
	ob := keysNeeded.OneBits()
	for _, i := range ob {
		key := rune('a' + i)
		if v.keysToUnlock[fromKey][key].AndEq(openDoors) {
			ret = append(ret, key)
		}
	}
	return
}

func (v *Vault) shortestPathFrom(c rune, keysNeeded, openDoors intmath.Bitmap) int {
	if keysNeeded.Val() == 0 {
		return 0
	}
	cacheKey := fmt.Sprintf("%v|%v", c, toLetters(keysNeeded))
	result, ok := v.score[cacheKey]
	if ok {
		return result
	}

	result = math.MaxInt32
	vc := v.visibleKeys(c, keysNeeded, openDoors)
	for _, key := range vc {
		newKeysNeeded := keysNeeded.Unset(int(key - 'a'))
		newOpenDoors := openDoors.Set(int(key - 'a'))
		d := v.distances[c][key] + v.shortestPathFrom(key, newKeysNeeded, newOpenDoors)
		if d < result {
			result = d
		}
	}
	v.score[cacheKey] = result
	return result
}

func (v *Vault) doorsWithoutKeys() (ret intmath.Bitmap) {
	for r := 'a'; r <= 'z'; r++ {
		if _, ok := v.keys[r]; !ok {
			ret = ret.Set(int(r - 'a'))
		}
	}
	return
}

// ShortestPath finds the shortest path from the start to all keys.
func (v *Vault) ShortestPath() int {
	v.score = map[string]int{}
	keysNeeded := intmath.Bitmap{}
	for key := range v.keys {
		keysNeeded = keysNeeded.Set((int(key - 'a')))
	}
	return v.shortestPathFrom('@', keysNeeded, v.doorsWithoutKeys())
}

// New creates a new Vault object and identifies key points.
func New(p string) *Vault {
	v := Vault{}
	v.m = input(p)
	v.findKeyPoints()
	v.findPaths()
	return &v
}

func input(p string) (ret [][]rune) {
	lines := fileinput.ReadLines(p)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		l := []rune{}
		for _, r := range line {
			l = append(l, r)
		}
		ret = append(ret, l)
	}
	return
}
