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
