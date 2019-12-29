package intmath

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
