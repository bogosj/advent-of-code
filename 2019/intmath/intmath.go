package intmath

// Point is a point in Cartesian space.
type Point struct {
	x, y int
}

// Gcd returns the greatest common denominator of two numbers.
// https://play.golang.org/p/SmzvkDjYlb
func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Lcm returns the least common multiple of two numbers.
func Lcm(a, b int, integers ...int) int {
	result := a * b / Gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = Lcm(result, integers[i])
	}

	return result
}

// Abs returns the absolute value of a number.
func Abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}
