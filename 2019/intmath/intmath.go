package intmath

// Point is a point in Cartesian space.
type Point struct {
	X, Y int
}

// Neighbors returns the neighboring points in horizontal and vertical directions.
func (p Point) Neighbors() (ret []Point) {
	ret = append(ret, Point{p.X - 1, p.Y})
	ret = append(ret, Point{p.X + 1, p.Y})
	ret = append(ret, Point{p.X, p.Y - 1})
	ret = append(ret, Point{p.X, p.Y + 1})
	return
}

// Bitmap provides a simple implementation of a 64 bit bitmap
type Bitmap struct {
	b uint64
}

// String returns a string representation of the Bitmap.
func (b Bitmap) String() string {
	return string(b.b)
}

// Unset sets the n-th bit to 0.
func (b Bitmap) Unset(i int) Bitmap {
	b.b &^= 1 << i
	return b
}

// Set sets the n-th bit to 1.
func (b Bitmap) Set(i int) Bitmap {
	b.b |= 1 << i
	return b
}

// OneBits returns the cardinal value of all bits set to 1.
func (b Bitmap) OneBits() (ret []int) {
	for i := 0; i < 64; i++ {
		if b.b&1 == 1 {
			ret = append(ret, i)
		}
		b.b = b.b >> 1
	}
	return
}

// AndEq bitwise ANDs the values and returns true if the result matches the called Bitmap.
func (b Bitmap) AndEq(o Bitmap) bool {
	return b.b&o.b == b.b
}

// Val exposes the underlying int value of the Bitmap.
func (b Bitmap) Val() uint64 {
	return b.b
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
