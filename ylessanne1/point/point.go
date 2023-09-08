package point

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) String() string {
	return fmt.Sprintf("x: %f\ny: %f\n", p.x, p.y)
}

func (p *Point) Roh() float64 {
	return math.Sqrt(math.Pow(float64(p.x), 2) + math.Pow(p.y, 2))
}

func (p *Point) Theta() float64 {
	return math.Atan2(p.y, p.x)
}
