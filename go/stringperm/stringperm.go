package stringperm

import "github.com/bogosj/advent-of-code/intmath"

// Permutations returns a channel with all permutations of a given slice of strings.
func Permutations(s []string) <-chan []string {
	out := make(chan []string)
	go func(o chan []string) {
		defer close(o)
		var idx []int
		for i := range s {
			idx = append(idx, i)
		}
		idxs := intmath.Permutations(idx)
		for p := range idxs {
			var ss []string
			for _, i := range p {
				ss = append(ss, s[i])
			}
			o <- ss
		}
	}(out)
	return out
}
