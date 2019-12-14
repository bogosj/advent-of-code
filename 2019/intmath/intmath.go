package intmath

// https://play.golang.org/p/SmzvkDjYlb
func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Lcm(a, b int, integers ...int) int {
	result := a * b / Gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = Lcm(result, integers[i])
	}

	return result
}

func Abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}
