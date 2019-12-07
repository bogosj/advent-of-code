package main

func permutations(in []int) (p [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			p = append(p, append([]int{}, a...))
		} else {
			for i := k; i < len(in); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(in, 0)
	return p
}
