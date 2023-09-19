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
	rho   float64
	theta float64
}

func NewPoint(x, y float64) *Point {
	r := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	t := math.Atan2(y, x)
	return &Point{rho: r, theta: t}
}

func (p *Point) String() string {
	return fmt.Sprintf("x: %f\ny: %f\nrho: %f\ntheta: %f\n", p.X(), p.Y(), p.Roh(), p.Theta())
}

func (p *Point) X() float64 {
	return p.Roh() * math.Cos(p.Theta())
}

func (p *Point) Y() float64 {
	return p.Roh() * math.Sin(p.Theta())
}

func (p *Point) Roh() float64 {
	return p.rho
}

func (p *Point) Theta() float64 {
	return p.theta
}

func (p *Point) Distance(pnt *Point) float64 {
	v := p.VectorTo(*pnt)
	return v.Roh()
}

// VectorTo returns point
func (p *Point) VectorTo(pnt Point) Point {
	return *NewPoint(pnt.X()-p.X(), pnt.Y()-p.Y())
}

/*
Translate moves point by given amount. dx horisontally and dy vertically.
POST: P(x, y).Translate(dx, dy) -> P(x+dx, y+dy)
*/
func (p *Point) Translate(dx float64, dy float64) {
	x := p.X() + dx
	y := p.Y() + dy

	p.rho = math.Sqrt(math.Pow(x, 2) + math.Sqrt(math.Pow(y, 2)))
	p.theta = math.Atan2(y, x)
}

/*
Scale by factor
POST: X() = X()*factor && Y() = Y()*factor
*/
func (p *Point) Scale(factor float64) {
	p.rho *= factor
}

/*
Rotate around origin (0.0, 0.0) by angle
POST: Theta() = Theta()+angle
*/
func (p *Point) CentreRotate(angle float64) {
	p.theta = p.Theta() + angle
}

/*
Rotate around point by angle
POST: math.Mod(p.theta, 2*math.Pi) = math.Mod((p.theta + angle), 2*math.Pi)
*/
func (p *Point) Rotate(pnt Point, angle float64) {
	p.Translate(-pnt.X(), -pnt.Y())
	p.CentreRotate(angle)
	p.Translate(pnt.X(), pnt.Y())
}
