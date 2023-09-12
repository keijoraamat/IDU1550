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

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) String() string {
	return fmt.Sprintf("x: %f\ny: %f\n", p.x, p.y)
}

func (p *Point) Roh() float64 {
	//Distance to origin (0.0, 0.0)
	return math.Sqrt(math.Pow(float64(p.x), 2) + math.Pow(p.y, 2))
}

func (p *Point) Theta() float64 {
	//Angle to horisontal axis
	return math.Atan2(p.y, p.x)
}

func (p *Point) Distance(other *Point) float64 {
	//Distance to other
	v := p.VectorTo(*other)
	return v.Roh()
}

func (p *Point) VectorTo(other Point) Point {
	//Point representing the vector from selt to other point
	return *NewPoint(other.x-p.x, other.y-p.y)
}

func (p *Point) Translate(dx float64, dy float64) {
	p.x += dx
	p.y += dy
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}

func (p *Point) CentreRotate(angle float64) {
	p.x = p.Roh() * math.Cos(p.Theta()+angle)
	p.y = p.Roh() * math.Sin(p.Theta()+angle)
}

func (p *Point) Rotate(pnt Point, angle float64) {
	p.Translate(-pnt.X(), -pnt.Y())
	p.CentreRotate(angle)
	p.Translate(pnt.X(), pnt.Y())
}
