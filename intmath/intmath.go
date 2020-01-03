package intmath

import (
	"math"
	"strconv"
)

// Atoi returns the integer value of a string or panics.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
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

// Min returns the minimum value provided.
func Min(in ...int) int {
	i := in[0]
	for _, v := range in {
		if v < i {
			i = v
		}
	}
	return i
}

// Max returns the minimum value provided.
func Max(in ...int) int {
	i := in[0]
	for _, v := range in {
		if v > i {
			i = v
		}
	}
	return i
}

// Sqrt returns floor(sqrt(i)).
func Sqrt(i int) int {
	f := math.Sqrt(float64(i))
	return int(f)
}
