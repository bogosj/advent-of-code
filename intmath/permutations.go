package intmath

// Permutations returns all of the permutations of a given slice of ints.
func Permutations(in []int) <-chan []int {
	out := make(chan []int)

	go func(o chan []int) {
		var rc func([]int, int)
		rc = func(a []int, k int) {
			if k == len(a) {
				o <- append([]int{}, a...)
			} else {
				for i := k; i < len(in); i++ {
					a[k], a[i] = a[i], a[k]
					rc(a, k+1)
					a[k], a[i] = a[i], a[k]
				}
			}
		}
		rc(in, 0)
		close(o)
	}(out)
	return out
}
