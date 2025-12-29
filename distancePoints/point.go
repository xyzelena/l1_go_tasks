package distancePoints

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p Point) Distance(q Point) float64 {
	result := math.Round(math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2)))
	return result
}
