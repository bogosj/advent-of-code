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

// AllNeighbors returns the neighboring points in all directions.
func (p Point) AllNeighbors() (ret []Point) {
	ret = p.Neighbors()
	ret = append(ret, Point{p.X - 1, p.Y - 1})
	ret = append(ret, Point{p.X + 1, p.Y + 1})
	ret = append(ret, Point{p.X + 1, p.Y - 1})
	ret = append(ret, Point{p.X - 1, p.Y + 1})
	return
}

// ManhattanDistanceTo computes the distance between two points on a grid.
func (p Point) ManhattanDistanceTo(op Point) int {
	return Abs(p.X-op.X) + Abs(p.Y-op.Y)
}
