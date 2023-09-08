package point

import "fmt"

type Point struct {
	x float32
	y float32
}

func NewPoint(x, y float32) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) String() string {
	return fmt.Sprintf("x: %f\ny: %f\n", p.x, p.y)
}
