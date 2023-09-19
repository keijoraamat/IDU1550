// Package point provides a simple representation of a two-dimensional point
// and utility functions for manipulating points in a Cartesian coordinate
// system.
//
// Points can be created with the NewPoint function, and various operations
// like translation, scaling, rotation, and distance calculation are provided.
// Additionally, this package includes methods for accessing the X and Y
// coordinates, calculating the distance to the origin, and finding the angle
// with respect to the horizontal axis.
//
// Example usage:
//
//	p := point.NewPoint(3.0, 4.0)
//	p.Translate(1.0, 2.0)
//	p.Scale(2.0)
//	p.Rotate(point.NewPoint(0.0, 0.0), math.Pi/2)
//
//	fmt.Printf("X: %f, Y: %f\n", p.X(), p.Y())
//	fmt.Printf("Distance to origin: %f\n", p.Roh())
//	fmt.Printf("Angle with horizontal axis: %f\n", p.Theta())
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
