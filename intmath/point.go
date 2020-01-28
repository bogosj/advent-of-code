package intmath

// Point is a point in Cartesian space.
type Point struct {
	X, Y int
}

// Neighbor returns a new point in a given directon. Valid directions
// are UDLR or NSEW (case insensitive).
func (p Point) Neighbor(dir rune) (op Point) {
	op.X, op.Y = p.X, p.Y
	switch dir {
	case 'U', 'u', 'N', 'n':
		op.Y--
	case 'D', 'd', 'S', 's':
		op.Y++
	case 'L', 'l', 'W', 'w':
		op.X--
	case 'R', 'r', 'E', 'e':
		op.X++
	}
	return
}

// Neighbors returns the neighboring points in horizontal and vertical directions.
// This function returns points in reading order.
func (p Point) Neighbors() (ret []Point) {
	ret = append(ret, Point{p.X, p.Y - 1})
	ret = append(ret, Point{p.X - 1, p.Y})
	ret = append(ret, Point{p.X + 1, p.Y})
	ret = append(ret, Point{p.X, p.Y + 1})
	return
}

// AllNeighbors returns the neighboring points in all directions.
// This function returns points in reading order.
func (p Point) AllNeighbors() (ret []Point) {
	ret = append(ret, Point{p.X - 1, p.Y - 1})
	ret = append(ret, Point{p.X, p.Y - 1})
	ret = append(ret, Point{p.X + 1, p.Y - 1})
	ret = append(ret, Point{p.X - 1, p.Y})
	ret = append(ret, Point{p.X + 1, p.Y})
	ret = append(ret, Point{p.X - 1, p.Y + 1})
	ret = append(ret, Point{p.X, p.Y + 1})
	ret = append(ret, Point{p.X + 1, p.Y + 1})
	return
}

// ManhattanDistanceTo computes the distance between two points on a grid.
func (p Point) ManhattanDistanceTo(op Point) int {
	return Abs(p.X-op.X) + Abs(p.Y-op.Y)
}
