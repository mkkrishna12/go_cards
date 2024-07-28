package main

type shape interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func (t square) getArea() float64 {
	return t.side * t.side
}
func (s triangle) getArea() float64 {
	return (0.5 * s.base * s.height)
}
func getTheAreaForShape(s shape) float64 {
	return s.getArea()
}
