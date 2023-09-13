package point

import "math"

type Point struct {
	rho   float64
	theta float64
}

func NewPoint(x, y float64) *Point {
	r := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	t := math.Atan2(y, x)
	return &Point{rho: r, theta: t}
}
